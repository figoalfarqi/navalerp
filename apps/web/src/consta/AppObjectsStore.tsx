export enum APP_OBJECT_STORE {
  DELIVERY_ORDER_ACTIVE = "delivery_order_active",
  DELIVERY_ORDER_ASSIGN_CHECK = "delivery_order_assign_check",
  HELLO_CHECK = "hello_check",
}


export const AppObjectStoreEndpointMap: Record<APP_OBJECT_STORE, string> = {
  [APP_OBJECT_STORE.DELIVERY_ORDER_ACTIVE]: "/driver/delivery_order/active",
  [APP_OBJECT_STORE.DELIVERY_ORDER_ASSIGN_CHECK]: "/driver/delivery_order/assign_check",
  [APP_OBJECT_STORE.HELLO_CHECK]: "/hello",
};