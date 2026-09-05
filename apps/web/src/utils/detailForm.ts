import { addMinutes } from "./dateTime";

// export const handleCascadeArrivalTime = ({
//     prevItem,
//     newItem,
//     updatedGroup,
//     others,
//     completed_at,
//     DEFAULT_INTERVAL_MINUTES,
// }: {
//     prevItem: any;
//     newItem: any;
//     updatedGroup: any[];
//     others: any[];
//     completed_at: any;
//     DEFAULT_INTERVAL_MINUTES: number;
// }) => {
//     let localUpdatedGroup = [...updatedGroup];
//     let localOthers = [...others];
//     let localCompletedAt = completed_at;

//     if (
//         prevItem.project_detail_destination_sequence === 1 &&
//         prevItem.project_detail_destination_type_id === 1 &&
//         newItem.arrival_time &&
//         typeof newItem.arrival_time === "string" &&
//         newItem.arrival_time !== prevItem.arrival_time
//     ) {
//         const secondSequencePickup = localUpdatedGroup.find(
//             (pickup) => pickup.project_detail_destination_sequence === 2,
//         );
//         const firstSequenceDrop = localOthers.find(
//             (drop) => drop.project_detail_destination_sequence === 1,
//         );

//         let lastAddedTime = new Date(newItem.arrival_time);
//         let lastToNextMinutes =
//             newItem.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

//         if (secondSequencePickup) {
//             const newSecondSequencePickup = {
//                 ...secondSequencePickup,
//                 arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
//             };
//             lastAddedTime = new Date(newSecondSequencePickup.arrival_time);
//             lastToNextMinutes =
//                 newSecondSequencePickup.to_next_duration_minutes ??
//                 DEFAULT_INTERVAL_MINUTES;

//             localUpdatedGroup = localUpdatedGroup.map((item) =>
//                 item.project_detail_destination_sequence === 2
//                     ? newSecondSequencePickup
//                     : item,
//             );

//             const thirdSequencePickup = localUpdatedGroup.find(
//                 (pickup) => pickup.project_detail_destination_sequence === 3,
//             );

//             if (thirdSequencePickup) {
//                 const newThirdSequencePickup = {
//                     ...thirdSequencePickup,
//                     arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
//                 };
//                 lastAddedTime = new Date(newThirdSequencePickup.arrival_time);
//                 lastToNextMinutes =
//                     newThirdSequencePickup.to_next_duration_minutes ??
//                     DEFAULT_INTERVAL_MINUTES;

//                 localUpdatedGroup = localUpdatedGroup.map((item) =>
//                     item.project_detail_destination_sequence === 3
//                         ? newThirdSequencePickup
//                         : item,
//                 );
//             }
//         }
//         if (firstSequenceDrop) {
//             const newFirstSequenceDrop = {
//                 ...firstSequenceDrop,
//                 arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
//             };
//             lastAddedTime = new Date(newFirstSequenceDrop.arrival_time);
//             lastToNextMinutes =
//                 newFirstSequenceDrop.to_next_duration_minutes ??
//                 DEFAULT_INTERVAL_MINUTES;

//             localOthers = localOthers.map((item) =>
//                 item.project_detail_destination_sequence === 1
//                     ? newFirstSequenceDrop
//                     : item,
//             );

//             const secondSequenceDrop = localOthers.find(
//                 (pickup) => pickup.project_detail_destination_sequence === 2,
//             );

//             if (secondSequenceDrop) {
//                 const newSecondSequenceDrop = {
//                     ...secondSequenceDrop,
//                     arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
//                 };
//                 lastAddedTime = new Date(newSecondSequenceDrop.arrival_time);
//                 lastToNextMinutes =
//                     newSecondSequenceDrop.to_next_duration_minutes ??
//                     DEFAULT_INTERVAL_MINUTES;

//                 localOthers = localOthers.map((item) =>
//                     item.project_detail_destination_sequence === 2
//                         ? newSecondSequenceDrop
//                         : item,
//                 );

//                 const thirdSequenceDrop = localOthers.find(
//                     (pickup) => pickup.project_detail_destination_sequence === 3,
//                 );

//                 if (thirdSequenceDrop) {
//                     const newThirdSequenceDrop = {
//                         ...thirdSequenceDrop,
//                         arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
//                     };
//                     lastAddedTime = new Date(newThirdSequenceDrop.arrival_time);
//                     lastToNextMinutes =
//                         newThirdSequenceDrop.to_next_duration_minutes ??
//                         DEFAULT_INTERVAL_MINUTES;
//                     localOthers = localOthers.map((item) =>
//                         item.project_detail_destination_sequence === 3
//                             ? newThirdSequenceDrop
//                             : item,
//                     );
//                 }
//             }
//         }

//         localCompletedAt = addMinutes(lastAddedTime, lastToNextMinutes);
//     }

//     return {
//         updatedGroup: localUpdatedGroup,
//         others: localOthers,
//         completed_at: localCompletedAt,
//     };
// };




const MAX_SEQUENCE = 10;

export const handleCascadeArrivalTime = ({
    prevItem,
    newItem,
    updatedGroup,
    others,
    completed_at,
    DEFAULT_INTERVAL_MINUTES,
}: {
    prevItem: any;
    newItem: any;
    updatedGroup: any[];
    others: any[];
    completed_at: any;
    DEFAULT_INTERVAL_MINUTES: number;
}) => {
    let localUpdatedGroup = [...updatedGroup];
    let localOthers = [...others];
    let localCompletedAt = completed_at;

    if (
        // prevItem.project_detail_destination_sequence === 1 &&
        prevItem.project_detail_destination_type_id === 1 &&
        newItem.arrival_time &&
        typeof newItem.arrival_time === "string" &&
        newItem.arrival_time !== prevItem.arrival_time
    ) {
        const nextDestinationSequence = prevItem.project_detail_destination_sequence + 1
        let lastAddedTime = new Date(newItem.arrival_time);
        let lastToNextMinutes =
            newItem.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

        // ===== PICKUP (updatedGroup) =====
        const hasPickupSequence2 = localUpdatedGroup.some(
            (i) => i.project_detail_destination_sequence === nextDestinationSequence,
        );

        if (hasPickupSequence2) {
            for (let seq = nextDestinationSequence; seq <= MAX_SEQUENCE; seq++) {
                const current = localUpdatedGroup.find(
                    (i) => i.project_detail_destination_sequence === seq,
                );
                if (!current) break;

                const updated = {
                    ...current,
                    arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
                };

                lastAddedTime = new Date(updated.arrival_time);
                lastToNextMinutes =
                    updated.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

                localUpdatedGroup = localUpdatedGroup.map((item) =>
                    item.project_detail_destination_sequence === seq ? updated : item,
                );
            }
        }

        // ===== DROP (others) =====
        for (let seq = 1; seq <= MAX_SEQUENCE; seq++) {
            const current = localOthers.find(
                (i) => i.project_detail_destination_sequence === seq,
            );
            if (!current) break;

            const updated = {
                ...current,
                arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
            };

            lastAddedTime = new Date(updated.arrival_time);
            lastToNextMinutes =
                updated.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

            localOthers = localOthers.map((item) =>
                item.project_detail_destination_sequence === seq ? updated : item,
            );
        }


        localCompletedAt = addMinutes(lastAddedTime, lastToNextMinutes);
    }
    else if (
        // prevItem.project_detail_destination_sequence === 1 &&
        prevItem.project_detail_destination_type_id === 2 &&
        newItem.arrival_time &&
        typeof newItem.arrival_time === "string" &&
        newItem.arrival_time !== prevItem.arrival_time
    ) {
        const nextDestinationSequence = prevItem.project_detail_destination_sequence + 1
        let lastAddedTime = new Date(newItem.arrival_time);
        let lastToNextMinutes =
            newItem.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

        // ===== DROP (updatedGroup) =====
        const hasDropSequence2 = localUpdatedGroup.some(
            (i) => i.project_detail_destination_sequence === nextDestinationSequence,
        );

        if (hasDropSequence2) {
            for (let seq = nextDestinationSequence; seq <= MAX_SEQUENCE; seq++) {
                const current = localUpdatedGroup.find(
                    (i) => i.project_detail_destination_sequence === seq,
                );
                if (!current) break;

                const updated = {
                    ...current,
                    arrival_time: addMinutes(lastAddedTime, lastToNextMinutes),
                };

                lastAddedTime = new Date(updated.arrival_time);
                lastToNextMinutes =
                    updated.to_next_duration_minutes ?? DEFAULT_INTERVAL_MINUTES;

                localUpdatedGroup = localUpdatedGroup.map((item) =>
                    item.project_detail_destination_sequence === seq ? updated : item,
                );
            }
        }

        localCompletedAt = addMinutes(lastAddedTime, lastToNextMinutes);
    }

    return {
        updatedGroup: localUpdatedGroup,
        others: localOthers,
        completed_at: localCompletedAt,
    };
};
