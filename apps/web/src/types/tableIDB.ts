export const STORES = {
    AUTH: "auth",
    IRREGULAR_COST: "irregular_cost",
    IRREGULAR_COST_TYPE: "irregular_cost_type",
    WORKSHOP: "workshop",
    CITY: "city",
    DISTRICT: "district",
    PENDING_POSTS: "pending_posts",
    LEAVE_REQUEST: "leave_request",
    LEAVE_REQUEST_TYPE: "leave_request_type",
    delivery_order_status: "delivery_order_status",
    customer_product: "customer_product",
    APP_OBJECTS: "app_objects",
    FORM_OBJECT: "form_objects",
    APP_FILES: "app_files",
} as const;

export type StoreName = (typeof STORES)[keyof typeof STORES];

export const DB_NAME = "pmlidb" 
export const DB_VERSION = 1