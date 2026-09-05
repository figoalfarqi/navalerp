"use client";

import {
  ButtonSize,
  ButtonVariant,
  SizeClassButtonMap,
  VariantClassesButtonMap,
} from "@/consta/VariantClassesButton";
import { ReactNode, useRef } from "react";

interface ButtonProps {
  id: string;
  type?: "button" | "submit" | "reset";
  variant?: ButtonVariant;
  size?: ButtonSize;
  className?: string;
  wrapperClassName?: string;
  onClick?: () => void;
  disabled?: boolean;
  children?: ReactNode;
}

export default function Button({
  id,
  type = "button",
  variant = "blue-solid",
  size = "md",
  className = "",
  wrapperClassName = "",
  onClick,
  disabled = false,
  children,
}: ButtonProps) {
  const baseWrapperClass = `inline-flex ${wrapperClassName}`;

  const baseButtonClass =
    "relative overflow-hidden cursor-pointer inline-flex items-center justify-center gap-2 font-medium rounded-md focus:outline-none transition disabled:opacity-50 disabled:cursor-not-allowed transition duration-200";

  const buttonRef = useRef<HTMLButtonElement>(null);

  const createRipple = (event: React.MouseEvent<HTMLButtonElement>) => {
    const button = buttonRef.current;
    if (!button) return;

    const ripple = document.createElement("span");
    const diameter = Math.max(button.clientWidth, button.clientHeight);
    const radius = diameter / 2;

    ripple.style.width = ripple.style.height = `${diameter}px`;
    ripple.style.left = `${
      event.clientX - button.getBoundingClientRect().left - radius
    }px`;
    ripple.style.top = `${
      event.clientY - button.getBoundingClientRect().top - radius
    }px`;
    ripple.classList.add("ripple");

    const existingRipple = button.getElementsByClassName("ripple")[0];
    if (existingRipple) existingRipple.remove();

    button.appendChild(ripple);

    // Remove ripple after animation
    setTimeout(() => ripple.remove(), 400);
  };

  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    // e.preventDefault();
    e.stopPropagation();
    createRipple(e);
    onClick?.();
  };

  return (
    <div className={baseWrapperClass}>
      <button
        id={id}
        type={type}
        ref={buttonRef}
        onClick={handleClick}
        disabled={disabled}
        className={`${baseButtonClass} ${SizeClassButtonMap[size]} ${VariantClassesButtonMap[variant]} ${className}`}
      >
        {children}
      </button>
    </div>
  );
}