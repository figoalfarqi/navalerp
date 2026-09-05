import MapTile from "./MapTile";

const CameraOverlay = ({
  location,
  address,
}: {
  location: { lat: number; lng: number };
  address: {
    alamat?: string;
    desa?: string;
    kecamatan?: string;
    kota?: string;
    provinsi?: string;
  };
}) => {
  return (
    <div className="absolute bottom-1 left-1 text-white leading-none pr-2 rounded-md  text-[10px] flex justify-between items-end w-full">
      <MapTile lat={location?.lat} lng={location?.lng} />
      <div className="space-y-1">
        <div>{new Date().toLocaleString()}</div>
        <div>Alamat: {address.alamat ?? "-"}</div>
        <div>Desa: {address.desa ?? "-"}</div>
        <div>Kecamatan: {address.kecamatan ?? "-"}</div>
        <div>Kota: {address.kota ?? "-"}</div>
        <div>Provinsi: {address.provinsi ?? "-"}</div>
      </div>
    </div>
  );
};
export default CameraOverlay;
