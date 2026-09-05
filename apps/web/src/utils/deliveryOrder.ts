
export function sortDeliveryOrderDestination(destinations: any) {
  return [...destinations].sort((a, b) => {
    if (a.delivery_order_destination_type_id !== b.delivery_order_destination_type_id) {
      return a.delivery_order_destination_type_id - b.delivery_order_destination_type_id;
    }

    return a.delivery_order_destination_sequence - b.delivery_order_destination_sequence;
  });
}

export function getLastDeliveryOrderStatusTypeIdStatusTime(destinations: any[]): {
    statusTypeId: number;
    statusTime: string;
    destinationName: string;
    districtName: string;
    cityName: string;
    provinceName: string;
    destinationIndex: number;
    statusIndex: number;
} {
    if (!Array.isArray(destinations) || destinations.length === 0)
        return {
            statusTypeId: -1,
            statusTime: "-",
            destinationName: "-",
            districtName: "-",
            cityName: "-",
            provinceName: "-",
            destinationIndex: 0,
            statusIndex: 0,
        };
    const orderedDestinations = sortDeliveryOrderDestination(destinations);

    for (let i = orderedDestinations.length - 1; i >= 0; i--) {
        const statuses = orderedDestinations[i]?.delivery_order_statuses;
        // console.log("AAAAAAAa",Array.isArray(statuses) && statuses.length > 0, statuses)
        if (Array.isArray(statuses) && statuses.length > 0) {
            return {
                statusTypeId:
                    statuses[statuses.length - 1].delivery_order_status_type_id,
                statusTime: statuses[statuses.length - 1].status_time,
                destinationName: orderedDestinations[i]?.destination.destination_name ?? "-",
                districtName:
                    orderedDestinations[i]?.destination.district?.district_name ?? "-",
                cityName: orderedDestinations[i]?.destination.city?.city_name ?? "-",
                provinceName:
                    orderedDestinations[i]?.destination.city?.province.province_name ?? "-",
                destinationIndex: i,
                statusIndex: statuses.length - 1,
            };
        }
    }

    return {
        statusTypeId: -1,
        statusTime: "-",
        destinationName: "-",
        districtName: "-",
        cityName: "-",
        provinceName: "-",
        destinationIndex: 0,
        statusIndex: 0,
    };
}

export function getNextDestination(
    destinations: any[],
    currentIndex: number
) {
    if (currentIndex + 1 >= destinations.length) return null;
    return destinations[currentIndex + 1];
}


export function resolveNextStatusType({
    lastStatusTypeID,
    hasNextPickup,
    hasNextDrop,
}: {
    lastStatusTypeID: number;
    hasNextPickup: boolean;
    hasNextDrop: boolean;
}): number | null {
    if (lastStatusTypeID === 5) {
        if (hasNextPickup) return 2;
        if (hasNextDrop) return 6;
        return null;
    }

    if (lastStatusTypeID === 9) {
        if (hasNextDrop) return 6;
        return 10;
    }

    return null;
}

export function getAutoNextStatus({
    currentStatusTypeID,
    nextDestinationType,
    hasNextDestination,
}: {
    currentStatusTypeID: number;
    nextDestinationType?: number;
    hasNextDestination: boolean;
}): number | null {
    if (currentStatusTypeID === 5) {
        if (!hasNextDestination) return null;
        return nextDestinationType === 0 ? 2 : 6;
    }

    if (currentStatusTypeID === 9) {
        return hasNextDestination ? 6 : null;
    }

    return null;
}
