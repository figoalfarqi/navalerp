const DocumentGuideOverlay = () => {
  return (
    <div className="absolute inset-0 pointer-events-none">
      {/* Layer gelap */}
      <div className="absolute inset-0 bg-black/20 rounded-sm" />

      {/* Area dokumen (hole) */}
      <div
        className="
          absolute
          top-1/2 left-1/2
          -translate-x-1/2 -translate-y-1/2
          w-[80%] h-[80%]
          border-2 border-white/80
          bg-white/20
          rounded-md
        "
      />

      {/* Text di tengah */}
      <div
        className="
          absolute
          top-1/2 left-1/2
          -translate-x-1/2 -translate-y-1/2
          text-white
          text-sm
          font-medium
          tracking-wide
        "
      >
        Foto dokumen di sini
      </div>
    </div>
  );
};

export default DocumentGuideOverlay;
