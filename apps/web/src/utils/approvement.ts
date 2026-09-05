import { AdminPayload } from "@/types/adminPayload";
import { formatDateTime } from "./dateTime";

export const isApproveableData = (
    {
        approvalData,
        adminPayload,
        firstStepAllowed,
        secondStepAllowed,
        thirdStepAllowed,
    }: {
        approvalData?: Record<string, any>[];
        adminPayload: AdminPayload | null;
        firstStepAllowed: number[];
        secondStepAllowed?: number[];
        thirdStepAllowed?: number[];
    }
) => {
    if (!adminPayload) {
        return { isApproveable: false, isApproveable1: false, isApproveable2: false, isApproveable3: false }
    }
    const isApproveable1 =
        (approvalData === undefined || approvalData?.length === 0 ||
            approvalData[approvalData.length - 1]
                .approval_step_id == 0) &&
        adminPayload &&
        firstStepAllowed.includes(adminPayload.app_role_id);
    const isApproveable2 = secondStepAllowed ?
        approvalData && approvalData?.length >= 1 &&
        approvalData[approvalData.length - 1]
            .approval_step_id == 1 &&
        adminPayload &&
        secondStepAllowed.includes(adminPayload.app_role_id) : false;

    const isApproveable3 = thirdStepAllowed ?
        approvalData && approvalData?.length >= 2 &&
        approvalData[approvalData.length - 1]
            .approval_step_id == 2 &&
        adminPayload &&
        thirdStepAllowed.includes(adminPayload.app_role_id) : false;

    return { isApproveable: isApproveable1 || isApproveable2 || isApproveable3, isApproveable1, isApproveable2, isApproveable3 }

}

export const countNotePhotosDOSByLastApprovalStatus = (
    {
        deliveryOrderStatuses,
        deliveryOrderStatus,
        validStatusIds,
        countWhenNoApproval = false,
    }: {
        deliveryOrderStatuses?: Record<string, any>[];
        deliveryOrderStatus?: Record<string, any>;
        validStatusIds: number[];
        countWhenNoApproval?: boolean;
    }): number => {
    if (!deliveryOrderStatuses && !deliveryOrderStatus) { return 0; }
    if (deliveryOrderStatuses) {
        return deliveryOrderStatuses
            .flatMap((dos: any) => dos?.delivery_order_note_photos ?? [])
            .reduce((acc: number, donp: any) => {
                const approvals = donp?.delivery_order_note_photo_approvals ?? [];

                // jika tidak ada approval
                if (approvals.length === 0) {
                    return countWhenNoApproval ? acc + 1 : acc;
                }

                const lastApproval = approvals.at(-1);

                if (lastApproval && validStatusIds.includes(lastApproval.approval_status_id)) {
                    return acc + 1;
                }

                return acc;
            }, 0);
    }
    else if (deliveryOrderStatus) {
        return (deliveryOrderStatus.delivery_order_note_photos ?? [])?.reduce((acc: number, donp: any) => {
            const approvals = donp?.delivery_order_note_photo_approvals ?? [];

            // jika tidak ada approval
            if (approvals.length === 0) {
                return countWhenNoApproval ? acc + 1 : acc;
            }

            const lastApproval = approvals.at(-1);

            if (lastApproval && validStatusIds.includes(lastApproval.approval_status_id)) {
                return acc + 1;
            }

            return acc;
        }, 0);
    }
    else { return 0; }
}




export const countCargoPhotosDOSByLastApprovalStatus = (
    {
        deliveryOrderStatuses,
        deliveryOrderStatus,
        validStatusIds,
        countWhenNoApproval = false,
    }: {
        deliveryOrderStatuses?: Record<string, any>[];
        deliveryOrderStatus?: Record<string, any>;
        validStatusIds: number[];
        countWhenNoApproval?: boolean;
    }): number => {
    if (!deliveryOrderStatuses && !deliveryOrderStatus) return 0;

    if (deliveryOrderStatuses) {
        return deliveryOrderStatuses
            .flatMap((dos: any) => dos?.delivery_order_cargo_photos ?? [])
            .reduce((acc: number, donp: any) => {
                const approvals = donp?.delivery_order_cargo_photo_approvals ?? [];

                // jika tidak ada approval
                if (approvals.length === 0) {
                    return countWhenNoApproval ? acc + 1 : acc;
                }

                const lastApproval = approvals.at(-1);

                if (lastApproval && validStatusIds.includes(lastApproval.approval_status_id)) {
                    return acc + 1;
                }

                return acc;
            }, 0);
    }
    else if (deliveryOrderStatus) {
        return (deliveryOrderStatus?.delivery_order_cargo_photos ?? [])
            .reduce((acc: number, donp: any) => {
                const approvals = donp?.delivery_order_cargo_photo_approvals ?? [];

                // jika tidak ada approval
                if (approvals.length === 0) {
                    return countWhenNoApproval ? acc + 1 : acc;
                }

                const lastApproval = approvals.at(-1);

                if (lastApproval && validStatusIds.includes(lastApproval.approval_status_id)) {
                    return acc + 1;
                }

                return acc;
            }, 0);
    }
    else {
        return 0;
    }
}

