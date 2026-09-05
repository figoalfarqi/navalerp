import { useEffect, useRef } from "react";

interface UseCloseOnScrollDistanceProps {
  triggerRef: React.RefObject<HTMLElement | null>;
  active: boolean;
  onClose: () => void;
  distance?: number; // default 200px
}

function getScrollParent(element: HTMLElement | null): HTMLElement | null {
  if (!element) return null;

  let parent = element.parentElement;

  while (parent) {
    const style = getComputedStyle(parent);
    const overflowY = style.overflowY;

    if (overflowY === "auto" || overflowY === "scroll") {
      return parent;
    }

    parent = parent.parentElement;
  }

  return document.documentElement;
}

// export function useCloseOnScrollDistance({
//   triggerRef,
//   active,
//   onClose,
//   distance = 200,
// }: UseCloseOnScrollDistanceProps) {
//   const initialScrollTopRef = useRef<number>(0);
//   const scrollParentRef = useRef<HTMLElement | null>(null);

//   useEffect(() => {
//     if (!active || !triggerRef.current) return;

//     const scrollParent = getScrollParent(triggerRef.current);
//     if (!scrollParent) return;

//     scrollParentRef.current = scrollParent;
//     initialScrollTopRef.current = scrollParent.scrollTop;

//     const handleScroll = () => {
//       const currentScroll = scrollParent.scrollTop;
//       const diff = Math.abs(currentScroll - initialScrollTopRef.current);

//       if (diff > distance) {
//         onClose();
//       }
//     };

//     scrollParent.addEventListener("scroll", handleScroll);

//     return () => {
//       scrollParent.removeEventListener("scroll", handleScroll);
//     };
//   }, [active, triggerRef, onClose, distance]);
// }


export function useCloseOnScrollDistance({
  triggerRef,
  active,
  onClose,
  distance = 200,
}: UseCloseOnScrollDistanceProps) {
  const lastScrollTopRef = useRef<number>(0);
  const accumulatedScrollRef = useRef<number>(0);
  const scrollParentRef = useRef<HTMLElement | null>(null);

  useEffect(() => {
    if (!active || !triggerRef.current) return;

    const scrollParent = getScrollParent(triggerRef.current);
    if (!scrollParent) return;

    scrollParentRef.current = scrollParent;
    lastScrollTopRef.current = scrollParent.scrollTop;
    accumulatedScrollRef.current = 0;

    const handleScroll = () => {
      const currentScroll = scrollParent.scrollTop;
      const delta = Math.abs(currentScroll - lastScrollTopRef.current);

      accumulatedScrollRef.current += delta;
      lastScrollTopRef.current = currentScroll;

      if (accumulatedScrollRef.current > distance) {
        onClose();
      }
    };

    scrollParent.addEventListener("scroll", handleScroll);

    return () => {
      scrollParent.removeEventListener("scroll", handleScroll);
    };
  }, [active, triggerRef, onClose, distance]);
}