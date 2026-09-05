const DesktopPageLoader = () => {
  return (
    <div className="flex items-center justify-center h-[80vh]">
      <div className="flex flex-col items-center space-y-2">
        <div className="h-10 w-10 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
        <p className="text-gray-600">Loading form...</p>
      </div>
    </div>
  );
};

export default DesktopPageLoader;
