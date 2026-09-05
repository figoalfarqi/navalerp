export async function waitForDBReady(getDbReady: () => boolean): Promise<void> {
  if (getDbReady()) return;

  await new Promise<void>((resolve) => {
    const interval = setInterval(() => {
      if (getDbReady()) {
        clearInterval(interval);
        resolve();
      }
    }, 50);
  });
}
