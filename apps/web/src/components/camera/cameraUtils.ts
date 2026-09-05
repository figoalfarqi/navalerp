let currentStream: MediaStream | null = null;

/**
 * Set media stream aktif
 */
export const setCameraStream = (stream: MediaStream | null) => {
  currentStream = stream;
};

/**
 * Stop semua track kamera yang aktif
 */
export const stopCamera = () => {
  if (!currentStream) return;
  currentStream.getTracks().forEach((track) => track.stop());
  currentStream = null;
};


type DrawOverlayParams = {
  ctx: CanvasRenderingContext2D;
  canvas: HTMLCanvasElement;
  location?: { lat: number; lng: number };
  address?: {
    alamat?: string;
    desa?: string;
    kecamatan?: string;
    kota?: string;
    provinsi?: string;
  };
  mapTileImage?: HTMLImageElement;
};

export function drawCameraOverlay({
  ctx,
  canvas,
  location,
  address,
  mapTileImage,
}: DrawOverlayParams) {
  const padding = 8;
  const tileSize = 110;
  const startY = canvas.height - tileSize - padding;
  const textX = canvas.width - padding;

  ctx.fillStyle = "#ffffff";
  ctx.font = "16px Arial";
  ctx.textAlign = "right";
  ctx.textBaseline = "top";

  let y = startY;

  ctx.fillText(new Date().toLocaleString(), textX, y);
  y += 20;

  ctx.fillText(`Alamat: ${address?.alamat ?? "-"}`, textX, y);
  y += 20;

  ctx.fillText(`Desa: ${address?.desa ?? "-"}`, textX, y);
  y += 20;

  ctx.fillText(`Kecamatan: ${address?.kecamatan ?? "-"}`, textX, y);
  y += 20;

  ctx.fillText(`Kota: ${address?.kota ?? "-"}`, textX, y);
  y += 20;

  ctx.fillText(`Provinsi: ${address?.provinsi ?? "-"}`, textX, y);

  // 🗺️ Draw map tile (kiri bawah)
  if (mapTileImage) {
    ctx.drawImage(
      mapTileImage,
      padding,
      startY,
      tileSize,
      tileSize
    );

    // marker
    ctx.font = "24px Arial";
    ctx.textAlign = "center";
    ctx.fillText(
      "📍",
      padding + tileSize / 2,
      startY + tileSize / 2
    );
  }
}


export function loadMapTile(lat: number, lng: number, zoom = 15): Promise<HTMLImageElement> {
  return new Promise((resolve, reject) => {
    const n = Math.pow(2, zoom);
    const latRad = (lat * Math.PI) / 180;

    const x = Math.floor(((lng + 180) / 360) * n);
    const y = Math.floor(
      ((1 - Math.log(Math.tan(latRad) + 1 / Math.cos(latRad)) / Math.PI) / 2) * n
    );

    const img = new Image();
    img.crossOrigin = "anonymous";
    img.src = `https://tile.openstreetmap.org/${zoom}/${x}/${y}.png`;

    img.onload = () => resolve(img);
    img.onerror = reject;
  });
}
