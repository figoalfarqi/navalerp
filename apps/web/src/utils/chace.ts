// src/utils/canNavigate.ts
export async function canNavigate(url: string): Promise<boolean> {
  if (!("caches" in window)) return false;

  const cacheNames = await caches.keys();

  for (const name of cacheNames) {
    const cache = await caches.open(name);
    const match = await cache.match(url);
    if (match) return true;
  }

  return false;
}
