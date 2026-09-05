/* eslint-disable @typescript-eslint/no-explicit-any */
import { useRef, useState, useEffect, useCallback } from "react";

type PointerGestureOptions = {
    /** Waktu threshold (ms) untuk dianggap long press */
    longPressThreshold?: number;
    /** Interval update progress (ms) */
    progressInterval?: number;
    /** Callback ketika long press selesai */
    onLongPress?: (index: number) => void;
    /** Callback ketika klik biasa (short press) */
    onClick?: (index: number) => void;
    /** Callback jika swipe berhasil */
    onSwipeSuccess?: (index: number) => void;
    /** Threshold jarak (px) untuk dianggap swipe sukses */
    swipeThreshold?: number;
};

export function usePointerGesture({
    longPressThreshold = 1000,
    progressInterval = 50,
    onLongPress,
    onClick,
    onSwipeSuccess,
    swipeThreshold = 120,
}: PointerGestureOptions = {}) {
    const timerRef = useRef<NodeJS.Timeout | null>(null);
    const longPressTriggered = useRef(false);
    const startTimeRef = useRef<number | null>(null);

    const startXRef = useRef<number | null>(null);
    const [swipeProgress, setSwipeProgress] = useState(0);

    const [progressMap, setProgressMap] = useState<{ [key: number]: number }>({});

    // 🟢 Tekan (mouseDown/touchStart)
    const handlePressStart = useCallback((idx: number, clientX?: number) => {
        longPressTriggered.current = false;
        startTimeRef.current = Date.now();
        startXRef.current = clientX ?? null;
        setSwipeProgress(0);

        timerRef.current = setInterval(() => {
            if (longPressTriggered.current) return;
            if (!startTimeRef.current) return;
            const elapsed = Date.now() - startTimeRef.current;
            const percent = Math.min((elapsed / longPressThreshold) * 100, 100);

            setProgressMap((prev) => ({
                ...prev,
                [idx]: percent,
            }));

            if (elapsed >= longPressThreshold) {
                clearInterval(timerRef.current!);
                timerRef.current = null;
                longPressTriggered.current = true;
                onLongPress?.(idx);
            }
        }, progressInterval);
    }, [longPressThreshold, progressInterval, onLongPress]);

    // 🟡 Gerakan drag (untuk swipe)
    const handleMove = useCallback(
        (idx: number, clientX?: number) => {
            if (!startXRef.current || !clientX) return;
            const deltaX = clientX - startXRef.current;
            const percent = Math.max(0, Math.min((deltaX / swipeThreshold) * 100, 100));
            setSwipeProgress(percent);

            if (percent >= 100) {
                onSwipeSuccess?.(idx);
                startXRef.current = null;
                setSwipeProgress(100);
            }
        },
        [swipeThreshold, onSwipeSuccess]
    );


    // 🔴 Lepas (mouseUp/touchEnd)
    const handlePressEnd = useCallback((idx: number) => {
        if (timerRef.current) clearInterval(timerRef.current);

        if (!longPressTriggered.current) {
            onClick?.(idx);
        }

        // reset progress
        setProgressMap((prev) => ({
            ...prev,
            [idx]: 0,
        }));

        setSwipeProgress(0);
        startTimeRef.current = null;
        startXRef.current = null;
    }, [onClick]);

    // Hentikan interval jika komponen unmount
    useEffect(() => {
        return () => {
            if (timerRef.current) clearInterval(timerRef.current);
        };
    }, []);

    return {
        /** Persentase progress long press (per index) */
        progressMap,
        setProgressMap,
        swipeProgress,
        setSwipeProgress,
        /** Event handler untuk memulai tekan */
        handlePressStart,
        handleMove,
        /** Event handler untuk melepas tekan */
        handlePressEnd,
    };
}
