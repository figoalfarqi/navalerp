"use client";
import { createContext, useContext, useState, ReactNode, useRef } from "react";
import { BiCheckCircle, BiErrorCircle } from "react-icons/bi";

type ToastType = "success" | "error";

interface Toast {
  id: number;
  type: ToastType;
  message: string;
}

interface ToastContextType {
  showToast: (duration: number, type: ToastType, message: string) => void;
  setWrapperClassname: (duration?: number, cls?: string) => void;
}

const ToastContext = createContext<ToastContextType | undefined>(undefined);

export const ToastProvider = ({ children }: { children: ReactNode }) => {
  const [toasts, setToasts] = useState<Toast[]>([]);
  const [wrapperClassname, setWrapperClassnameState] = useState<string>();

  // 🔹 refs untuk kontrol timeout
  const wrapperTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const wrapperExpireAtRef = useRef<number>(0);

  const showToast = (duration: number, type: ToastType, message: string) => {
    const id = Date.now();
    setToasts((prev) => [...prev, { id, type, message }]);

    setTimeout(() => {
      setToasts((prev) => prev.filter((t) => t.id !== id));
    }, duration);
  };

  const setWrapperClassname = (duration = 3000, cls?: string) => {
    if (!cls) {
      setWrapperClassnameState(undefined);
      return;
    }

    const now = Date.now();
    const newExpireAt = now + duration;

    // ambil waktu TERLAMA
    if (newExpireAt > wrapperExpireAtRef.current) {
      wrapperExpireAtRef.current = newExpireAt;
    }

    setWrapperClassnameState(cls);

    // reset timeout lama
    if (wrapperTimeoutRef.current) {
      clearTimeout(wrapperTimeoutRef.current);
    }

    wrapperTimeoutRef.current = setTimeout(() => {
      setWrapperClassnameState(undefined);
      wrapperExpireAtRef.current = 0;
    }, wrapperExpireAtRef.current - now);
  };

  const baseClass = "fixed top-5 right-5 flex flex-col gap-2 z-50";

  return (
    <ToastContext.Provider value={{ showToast, setWrapperClassname }}>
      {children}

      <div
        className={
          wrapperClassname ? `${baseClass} ${wrapperClassname}` : baseClass
        }
      >
        {toasts.map((toast) => (
          <div
            key={toast.id}
            className={`flex items-center gap-2 p-3 rounded shadow-md text-white ${
              toast.type === "success" ? "bg-green-500" : "bg-red-500"
            }`}
          >
            {toast.type === "success" ? (
              <BiCheckCircle size={20} />
            ) : (
              <BiErrorCircle size={20} />
            )}
            <span>{toast.message}</span>
          </div>
        ))}
      </div>
    </ToastContext.Provider>
  );
};

export const useToast = () => {
  const context = useContext(ToastContext);
  if (!context) {
    throw new Error("useToast must be used within a ToastProvider");
  }
  return context;
};
