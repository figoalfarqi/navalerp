// import { FaRegUser } from "react-icons/fa";
// import { MdCheck, MdClose, MdComputer } from "react-icons/md";
// const ApprovalMarkIcon = ({
//   approvalData,
//   onClick,
//   className,
//   isMultiStep,
// }: {
//   approvalData?: Record<string, any>[];
//   onClick?: () => void;
//   className?: string;
//   isMultiStep?: boolean;
// }) => {
//   return (
//     <>
//       {[1, 4].includes(approvalData?.[0]?.approval_status_id) ? (
//         <div
//           className={`${className} text-white bg-green-500/50 rounded-full`}
//           onClick={() => {
//             onClick?.()
//           }}
//         >
//           <MdCheck size={sizeMap[size]} />
//         </div>
//       ) : [0, 2, 5].includes(approvalData?.[0]?.approval_status_id) ? (
//         <div
//           className={`${className} text-white p-1 bg-yellow-500/50 rounded-full`}
//           onClick={() => {
//             onClick?.()
//           }}
//         >
//           <FaRegUser size={sizeMap[size]} />
//         </div>
//       ) : [3, 6].includes(approvalData?.[0]?.approval_status_id) ? (
//         <div
//           className={`${className} text-white bg-red-500/50 rounded-full`}
//           onClick={() => {
//             onClick?.()
//           }}
//         >
//           <MdClose size={sizeMap[size]} />
//         </div>
//       ) : (
//         <div
//           className={`${className} text-white p-0.75 bg-yellow-500/50 rounded-full`}
//           onClick={() => {
//             onClick?.()
//           }}
//         >
//           <MdComputer size={sizeMap[size]} />
//         </div>
//       )}
//     </>
//   );
// };
// export default ApprovalMarkIcon;

// perbaiki

// untuk isMultistep true atau false
// ketika tidak ada approvalData maka muncul <MdComputer size={sizeMap[size]} bg-yellow-500/50 />
// ketika approvalData index terakhir approval_step_id == 0 dan approval_status_id == 0 maka muncul <FaRegUser size={sizeMap[size]} bg-yellow-500/50 /> dan <LuClock2 size={sizeMap[size]} bg-yellow-500/50 />
// ketika approvalData index terakhir approval_step_id == 1 dan approval_status_id == 2 maka muncul <MdOutlineAssignmentLate size={sizeMap[size]} bg-yellow-500/50 />
// ketika approvalData index terakhir approval_step_id == 1 dan approval_status_id == 3 maka muncul <MdClose size={sizeMap[size]} bg-red-500/50 />

// ketika isMultiStep true

// ketika approvalData index terakhir approval_step_id == 1 dan approval_status_id == 1 maka tampil 2 icon yaitu <MdCheck size={sizeMap[size]} bg-green-500/50 /> dan <LuClock2 size={sizeMap[size]} bg-yellow-500/50 />
// ketika approvalData index sebelum terakhir approval_step_id == 1 dan approval_status_id == 1  dan index terakhir approval_step_id == 2 dan approval_status_id == 4 maka tampil 2 icon yaitu <MdCheck size={sizeMap[size]} bg-green-500/50 /> dan <MdCheck size={sizeMap[size]} bg-green-500/50 />
// ketika approvalData index sebelum terakhir approval_step_id == 1 dan approval_status_id == 1  dan index terakhir approval_step_id == 2 dan approval_status_id == 5 maka tampil 2 icon yaitu <MdCheck size={sizeMap[size]} bg-green-500/50 /> dan <MdOutlineAssignmentLate size={sizeMap[size]} bg-yellow-500/50 />
// ketika approvalData index sebelum terakhir approval_step_id == 1 dan approval_status_id == 1  dan index terakhir approval_step_id == 2 dan approval_status_id == 6 maka tampil 2 icon yaitu <MdCheck size={sizeMap[size]} bg-green-500/50 /> dan <MdClose size={sizeMap[size]} bg-red-500/50 />

// ketika isMultiStep false
// ketika approvalData index terakhir approval_step_id == 1 dan approval_status_id == 1 maka <MdCheck size={sizeMap[size]} bg-green-500/50 />

import { FaRegUser } from "react-icons/fa";
import {
  MdCheck,
  MdClose,
  MdComputer,
  MdOutlineAssignmentLate,
} from "react-icons/md";
import { LuClock2 } from "react-icons/lu";


const wrap = (
  className: string | undefined,
  bg: string,
  children: React.ReactNode,
  onClick?: () => void,
) => (
  <div
    className={`${className} ${bg} text-white rounded-full flex items-center gap-0.5 p-1`}
    onClick={onClick}
  >
    {children}
  </div>
);


type Props = {
  approvalData?: Record<string, any>[];
  onClick?: () => void;
  className?: string;
  isMultiStep?: boolean;
  size?:"sm" | "md" | "lg" | "xl"
};

const sizeMap = {sm : 14, md: 16, lg: 18, xl: 20}


const ApprovalMarkIcon = ({
  approvalData,
  onClick,
  className,
  isMultiStep = false,
  size="md"
}: Props) => {
  // ===============================
  // NO DATA
  // ===============================
  if (!approvalData || approvalData.length === 0) {
    return wrap(
      className,
      "bg-yellow-500/50",
      <MdComputer size={sizeMap[size]} />,
      onClick,
    );
  }

  const last = approvalData[approvalData.length - 1];
  const prev = approvalData[approvalData.length - 2];

  if (last.approval_step_id === 0 && last.approval_status_id === 0) {
    return wrap(
      className,
      "bg-yellow-500/50",
      <>
        <FaRegUser size={sizeMap[size]} />
        {/* <LuClock2 size={sizeMap[size]} /> */}
      </>,
      onClick,
    );
  }

  // ===============================
  // SINGLE STEP
  // ===============================
  if (!isMultiStep) {

    if (last.approval_step_id === 1 && last.approval_status_id === 2) {
      return wrap(
        className,
        "bg-yellow-500/50",
        <MdOutlineAssignmentLate size={sizeMap[size]} />,
        onClick,
      );
    }

    if (last.approval_step_id === 1 && last.approval_status_id === 3) {
      return wrap(className, "bg-red-500/50", <MdClose size={sizeMap[size]} />, onClick);
    }

    if (last.approval_step_id === 1 && last.approval_status_id === 1) {
      return wrap(className, "bg-green-500/50", <MdCheck size={sizeMap[size]} />, onClick);
    }
  }

  // ===============================
  // MULTI STEP
  // ===============================
  if (isMultiStep) {
    // step 1 approved, waiting next
    if (last.approval_step_id === 1 && last.approval_status_id === 1) {
      return (
        <>
          {wrap(
            className,
            "bg-green-500/50",
            <>
              <MdCheck size={sizeMap[size]} />
            </>,
            onClick,
          )}
          {wrap(
            className,
            "bg-yellow-500/50",
            <>
              <LuClock2 size={sizeMap[size]} />
            </>,
            onClick,
          )}
        </>
      );
    }

    if (
      prev?.approval_step_id === 1 &&
      prev?.approval_status_id === 1 &&
      last.approval_step_id === 2
    ) {
      if (last.approval_status_id === 4) {
        return (
          <>
            {wrap(
              className,
              "bg-green-500/50",
              <>
                <MdCheck size={sizeMap[size]} />
              </>,
              onClick,
            )}
            {wrap(
              className,
              "bg-green-500/50",
              <>
                <MdCheck size={sizeMap[size]} />
              </>,
              onClick,
            )}
          </>
        );
      }

      if (last.approval_status_id === 5) {
        return (
          <>
            {wrap(
              className,
              "bg-green-500/50",
              <>
                <MdCheck size={sizeMap[size]} />
              </>,
              onClick,
            )}
            {wrap(
              className,
              "bg-yellow-500/50",
              <>
                <MdOutlineAssignmentLate size={sizeMap[size]} />
              </>,
              onClick,
            )}
          </>
        );
      }

      if (last.approval_status_id === 6) {
        return (
          <>
            {wrap(
              className,
              "bg-green-500/50",
              <>
                <MdCheck size={sizeMap[size]} />
              </>,
              onClick,
            )}
            {wrap(
              className,
              "bg-red-500/50",
              <>
                <MdClose size={sizeMap[size]} />
              </>,
              onClick,
            )}
          </>
        );
      }
    }
  }

  // ===============================
  // FALLBACK
  // ===============================
  return wrap(className, "bg-yellow-500/50", <MdComputer size={sizeMap[size]} />, onClick);
};

export default ApprovalMarkIcon;
