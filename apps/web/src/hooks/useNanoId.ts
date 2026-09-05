import { useCallback } from "react";

export const useNanoId = () => {

    const generateId = useCallback(() => {
        return `${Date.now().toString(36)}_${Math.random()
            .toString(36)
            .slice(2, 8)}`;
    }, []);

    return { generateId };
};
