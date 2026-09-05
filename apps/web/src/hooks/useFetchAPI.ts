/* eslint-disable @typescript-eslint/no-explicit-any */
// hooks/useFetchAPI.ts
"use client";

import { useAuth } from "@/context/AuthContext";
import { useCallback, useMemo } from "react";

export interface APIResponse<T = any> {
    data: T | null;
    code: number;
    status: string;
    message: string;
    errors: any;
}

export type authTokenType = "driver" | "admin" | "checker" | "none";

export interface Options {
    authToken: authTokenType // "" | "driver" | custom string
    timeout?: number;   // default 10000 ms
    signal?: AbortSignal;
}

export function useFetchAPI() {
    const { driverToken, adminToken, checkerToken } = useAuth();

    const request = useCallback(async <T>(
        method: string,
        url: string,
        options: Options,
        body?: any
    ): Promise<APIResponse<T>> => {
        const controller = new AbortController();
        const timeout = options.timeout ?? 10000;

        const signal = options.signal ?? controller.signal;
        const id = setTimeout(() => controller.abort(), timeout);

        const headers: Record<string, string> = {};

        // Setup Authorization
        if (options.authToken !== "none") {
            const token =
                options.authToken === "driver"
                    ? driverToken
                    : options.authToken === "admin"
                      ? adminToken
                      : checkerToken;
            if (token) headers["Authorization"] = `Bearer ${token}`;
        }

        // console.log(options.authToken, adminToken, headers["Authorization"])

        // Setup Content-Type (skip kalau FormData)
        if (!(body instanceof FormData)) {
            headers["Content-Type"] = "application/json";
        }

        try {
            const response = await fetch(url, {
                method,
                headers,
                body: body
                    ? body instanceof FormData
                        ? body
                        : JSON.stringify(body)
                    : undefined,
                signal: signal,
            });

            clearTimeout(id);

            let data: any = null;
            let message = "";
            let errors: any = null;

            try {
                const json = await response.json();
                data = json.data ?? null;
                message = json.message ?? "";
                errors = json.errors ?? null;
            } catch {
                data = null;
            }

            // Global Error Handling: unauthorized / forbidden
            if (response.status === 401 || response.status === 403) {
                // Clear context & IndexedDB token
               // next TODO
                // clearDriverToken();
                // clearAdminToken();
            }

            return {
                data,
                code: response.status,
                status: response.statusText,
                message,
                errors,
            };
        } catch (error: any) {
            clearTimeout(id);

            if (error.name === "AbortError") {
                return {
                    data: null,
                    code: 408,
                    status: "Request Timeout",
                    message: "Request aborted by timeout",
                    errors: error,
                };
            }

            return {
                data: null,
                code: 500,
                status: "Network Error",
                message: error.message,
                errors: error,
            };
        }
    }, [adminToken, checkerToken, driverToken]);

    return useMemo(() => ({
        getAPI: <T>(url: string, options: Options) =>
            request<T>("GET", url, options),
        postAPI: <T>(url: string, options: Options, body?: any) =>
            request<T>("POST", url, options, body),
        putAPI: <T>(url: string, options: Options, body?: any) =>
            request<T>("PUT", url, options, body),
        patchAPI: <T>(url: string, options: Options, body?: any) =>
            request<T>("PATCH", url, options, body),
        deleteAPI: <T>(url: string, options: Options, body?: any) =>
            request<T>("DELETE", url, options, body),
        uploadAPI: <T>(url: string, options: Options, body?: FormData) =>
            request<T>("POST", url, options, body),
    }), [request]);
}
