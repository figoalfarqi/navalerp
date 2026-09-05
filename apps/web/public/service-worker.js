// ===============================
// SERVICE WORKER VERSION
// ===============================
const SW_VERSION = "v7"; // ⬅️ bump tiap deploy

const STATIC_CACHE = `static-${SW_VERSION}`;

// ===============================
// STATIC ASSETS (APP SHELL)
// ===============================
const ASSETS = [
  "/",
  "/favicon.ico",
  "/manifest.json",
  "/icons/icon-192x192.png",
  "/icons/icon-512x512.png",
];

// ===============================
// INSTALL
// ===============================
self.addEventListener("install", (event) => {
  console.log("[SW] Install", SW_VERSION);

  event.waitUntil(
    caches.open(STATIC_CACHE).then((cache) => cache.addAll(ASSETS)),
  );

  self.skipWaiting();
});

// ===============================
// ACTIVATE
// ===============================
self.addEventListener("activate", (event) => {
  console.log("[SW] Activate", SW_VERSION);

  event.waitUntil(
    caches.keys().then((keys) =>
      Promise.all(
        keys
          .filter((key) => key !== STATIC_CACHE)
          .map((key) => {
            console.log("[SW] Delete old cache:", key);
            return caches.delete(key);
          }),
      ),
    ),
  );

  self.clients.claim();
});

// ===============================
// FETCH
// ===============================
self.addEventListener("fetch", (event) => {
  const { request } = event;

  // ❌ Hanya handle GET
  if (request.method !== "GET") return;

  // ❌ Jangan ganggu SSE / streaming
  if (request.headers.get("accept") === "text/event-stream") return;

  // ===============================
  // API → NETWORK ONLY (PENTING)
  // ===============================
  if (request.url.includes("/api/")) {
    event.respondWith(fetch(request));
    return;
  }

  // ===============================
  // APP SHELL → CACHE FIRST
  // ===============================
  event.respondWith(
    caches.match(request).then((cached) => {
      // if (cached) return cached;

      return fetch(request).then((res) => {
        const clone = res.clone();
        caches.open(STATIC_CACHE).then((cache) => {
          cache.put(request, clone);
        });
        return res;
      });
    }),
  );
});

// ===============================
// PUSH NOTIFICATION
// ===============================
self.addEventListener("push", (event) => {
  console.log("[SW] Push received");

  const data = event.data?.json() || {};

  event.waitUntil(
    self.registration.showNotification(data.title || "Notification", {
      body: data.body || "",
      data,
    }),
  );
});

// ===============================
// MESSAGE
// ===============================
self.addEventListener("message", (event) => {
  if (event.data === "SKIP_WAITING") {
    self.skipWaiting();
  }
});
