"use client";

import { FaTruckMoving } from "@/components/icons";

type Props = {
  lat: number;
  lng: number;
  zoom?: number;
};

export default function MapTile({ lat, lng, zoom = 15 }: Props) {
  const { x, y } = latLngToTile(lat, lng, zoom);
  const tileUrl = `https://tile.openstreetmap.org/${zoom}/${x}/${y}.png`;

  return (
    <div className="relative w-16 h-16 rounded overflow-hidden">
      <img
        src={tileUrl}
        className="w-full h-full block"
        alt="map"
      />

      {/* Marker orang */}
      <div
        style={{
          position: "absolute",
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -100%)",
          fontSize: 10,
          pointerEvents: "none",
        }}
        aria-label="marker"
      >
        📍
      </div>
    </div>
  );
}

function latLngToTile(lat: number, lng: number, zoom: number) {
  const latRad = (lat * Math.PI) / 180;
  const n = Math.pow(2, zoom);

  const x = Math.floor(((lng + 180) / 360) * n);
  const y = Math.floor(
    ((1 - Math.log(Math.tan(latRad) + 1 / Math.cos(latRad)) / Math.PI) / 2) * n
  );

  return { x, y };
}
