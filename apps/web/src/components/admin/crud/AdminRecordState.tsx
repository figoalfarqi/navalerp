interface AdminRecordStateProps {
  error?: string;
}

export default function AdminRecordState({ error }: AdminRecordStateProps) {
  if (error) {
    return (
      <div className="rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
        {error}
      </div>
    );
  }

  return (
    <div className="grid min-h-64 place-items-center rounded-2xl border border-slate-200 bg-white">
      <div className="text-center">
        <div className="mx-auto h-9 w-9 animate-spin rounded-full border-4 border-blue-500 border-t-transparent" />
        <p className="mt-3 text-sm text-slate-500">Memuat data...</p>
      </div>
    </div>
  );
}
