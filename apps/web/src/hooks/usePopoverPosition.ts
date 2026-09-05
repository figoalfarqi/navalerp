// import { useState, useLayoutEffect } from "react";

// interface PopoverPosition {
//   top: number;
//   left: number;
// }

// type UsePopoverPositionProps = {
//   inputRef: React.RefObject<HTMLElement | null>;
//   visible: boolean;
//   offset?: number;
//   popoverHeight?: number;
// };

// export function usePopoverPosition({
//   inputRef,
//   visible,
//   offset = 8,
//   popoverHeight = 300,
// }: UsePopoverPositionProps) {
//   const [position, setPosition] = useState<PopoverPosition>({ top: 0, left: 0 });

//   useLayoutEffect(() => {
//     if (!visible || !inputRef.current) return;

//     const updatePosition = () => {
//       const rect = inputRef.current!.getBoundingClientRect();
//       const scrollY = window.scrollY;
//       const scrollX = window.scrollX;
//       const spaceBelow = window.innerHeight - rect.bottom;
//       const spaceAbove = rect.top;

//       let top: number;
//       // kalau tidak muat di bawah tapi muat di atas, naik ke atas input
//       if (spaceBelow < popoverHeight && spaceAbove > popoverHeight) {
//         top = rect.top + scrollY - popoverHeight - offset;
//       } else {
//         // kalau tidak muat di mana pun, boleh menutupi input
//         top = rect.bottom + scrollY + offset;
//         if (top + popoverHeight > window.innerHeight + scrollY) {
//           top = window.innerHeight + scrollY - popoverHeight - 4;
//         }
//       }

//       // Pastikan tidak keluar kiri/kanan layar
//       let left = rect.left + scrollX;
//       const maxLeft = window.innerWidth + scrollX - rect.width;
//       if (left + rect.width > maxLeft) {
//         left = Math.max(scrollX + 8, maxLeft - 8);
//       }

//       setPosition({ top, left });
//     };

//     updatePosition();
//     window.addEventListener("resize", updatePosition);
//     window.addEventListener("scroll", updatePosition, true);

//     return () => {
//       window.removeEventListener("resize", updatePosition);
//       window.removeEventListener("scroll", updatePosition, true);
//     };
//   }, [visible, inputRef, offset, popoverHeight]);

//   return position;
// }



import { useState, useLayoutEffect } from "react";

type UsePopoverPositionProps = {
  inputRef: React.RefObject<HTMLElement | null>;
  popoverRef?: React.RefObject<HTMLElement | null>;
  visible: boolean;
  offset?: number;
  popoverHeight?: number;
};


interface PopoverPosition {
  top: number;
  left: number;
}

export function usePopoverPosition({
  inputRef,
  popoverRef,
  visible,
  offset = 8,
  popoverHeight = 300,
}: UsePopoverPositionProps) {
  const [position, setPosition] = useState<PopoverPosition>({
    top: 0,
    left: 0,
  });

  useLayoutEffect(() => {
    if (!visible || !inputRef.current) return;

    const updatePosition = () => {
      const inputRect = inputRef.current!.getBoundingClientRect();
      const popoverRect = popoverRef?.current?.getBoundingClientRect();

      const scrollX = window.scrollX;
      const scrollY = window.scrollY;

      const viewportWidth = window.innerWidth;
      const viewportHeight = window.innerHeight;

      const popoverW = popoverRect?.width ?? inputRect.width;
      const popoverH = popoverRect?.height ?? popoverHeight;

      // =========================
      // VERTICAL (TOP)
      // =========================
      const spaceBelow = viewportHeight - inputRect.bottom;
      const spaceAbove = inputRect.top;

      let top: number;

      if (spaceBelow < popoverH && spaceAbove > popoverH) {
        // buka ke atas
        top = inputRect.top + scrollY - popoverH - offset;
      } else {
        // buka ke bawah
        top = inputRect.bottom + scrollY + offset;

        // clamp kalau kepanjangan
        if (top + popoverH > viewportHeight + scrollY) {
          top = viewportHeight + scrollY - popoverH - 4;
        }
      }

      // =========================
      // HORIZONTAL (LEFT)
      // =========================
      let left = inputRect.left + scrollX;

      // kalau overflow kanan → geser ke kiri
      if (left + popoverW > viewportWidth + scrollX) {
        left = viewportWidth + scrollX - popoverW - 8;
      }

      // kalau overflow kiri → clamp
      if (left < scrollX + 8) {
        left = scrollX + 8;
      }

      setPosition({ top, left });
    };

    updatePosition();

    window.addEventListener("resize", updatePosition);
    window.addEventListener("scroll", updatePosition, true);

    return () => {
      window.removeEventListener("resize", updatePosition);
      window.removeEventListener("scroll", updatePosition, true);
    };
  }, [visible, inputRef, popoverRef, offset, popoverHeight]);

  return position;
}
