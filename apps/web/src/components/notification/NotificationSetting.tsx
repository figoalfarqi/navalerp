"use client";

import { useEffect, useState } from "react";
import { usePushNotification } from "@/hooks/usePushNotification";
import Modal from "@/components/Modal";
import Button from "@/components/form/Button";
import TextField from "@/components/form/TextField";
import {
  FaDesktop,
  FaMobileAlt,
  FaCheckCircle,
  FaSignOutAlt,
  FaEdit,
  FaBell,
  FaBellSlash,
  FaTimesCircle,
} from "@/components/icons";
import LoaderDots from "@/components/form/LoaderDots";
import { authTokenType } from "@/hooks/useFetchAPI";

type Device = {
  device_id: string;
  device_label: string;
  browser: string;
  platform: string;
  is_mobile: boolean;
  user_agent: string;
  last_seen: string;
};

export default function NotificationSetting({
  authTokenType,
}: {
  authTokenType: authTokenType;
}) {
  const {
    isSupported,
    isSubscribed,
    setIsSubscribed,
    getDeviceId,
    subscribe,
    unsubscribeCurrentDevice,
    unsubscribeOtherDevice,
    getDevices,
    updateDeviceLabel,
    sendTestPush,
  } = usePushNotification("delivery_order", authTokenType);

  const [devices, setDevices] = useState<Device[]>([]);
  const [isShowModalDenied, setIsShowModalDenied] = useState(false);
  const [isShowModalUnsubscribeOther, setIsShowModalUnsubscribeOther] =
    useState(false);
  const [otherDevice, setOtherDevice] = useState<Device>();
  const [isLoadingDevices, setIsLoadingDevices] = useState(false);
  const [isLoadingSubscribe, setIsLoadingSubscribe] = useState(false);
  const [isLoadingPushTest, setIsLoadingPushTest] = useState(false);
  const [isEditing, setIsEditing] = useState<string | null>(null);
  const [labelInput, setLabelInput] = useState("");

  const currentDeviceId = getDeviceId();

  /* ---------------- load devices ---------------- */
  const loadDevices = async () => {
    setIsLoadingDevices(true);
    try {
      const res = await getDevices();
      setDevices(res.data || []);
    } finally {
      setIsLoadingDevices(false);
    }
  };

  useEffect(() => {
    loadDevices();
  }, []);

  useEffect(() => {
    if (!devices.length) {
      setIsSubscribed(false);
      return;
    }
    const isCurrentDeviceActive = devices.some(
      (d) => d.device_id === currentDeviceId,
    );

    setIsSubscribed(isCurrentDeviceActive);
  }, [devices, currentDeviceId, setIsSubscribed]);

  /* ---------------- handlers ---------------- */

  const handleSubscribe = async () => {
    setIsLoadingSubscribe(true);
    if (Notification.permission === "denied") {
      setIsShowModalDenied(true);
    } else {
      await subscribe();
    }

    await loadDevices();
    setIsLoadingSubscribe(false);
  };

  const handleUnsubscribeCurrentDevice = async () => {
    setIsLoadingSubscribe(true);
    await unsubscribeCurrentDevice();
    await loadDevices();
    setIsLoadingSubscribe(false);
  };

  const handleRename = async (deviceId: string) => {
    if (!labelInput.trim()) return;

    await updateDeviceLabel(deviceId, labelInput);
    setIsEditing(null);
    setLabelInput("");
    await loadDevices();
  };

  const handleUnsubscribeOther = async (deviceId: string) => {
    await unsubscribeOtherDevice(deviceId);
    await loadDevices();
  };

  const handleSendPushTest = async () => {
    setIsLoadingPushTest(true);
    await sendTestPush();
    setIsLoadingPushTest(false);
  };

  return (
    <div>
      {/* -------- status -------- */}
      <section className="mb-4 rounded-sm border border-gray-300 bg-white p-3 flex flex-col gap-2">
        {/* -------- status row -------- */}
        <div className="flex flex-col gap-1 text-base">
          <div className="flex items-center gap-2">
            {isSupported ? (
              <FaCheckCircle className="text-green-600" />
            ) : (
              <FaTimesCircle className="text-red-600" />
            )}
            <span className="font-medium">
              Browser Support:
              <span className="ml-1">
                {isSupported ? "Supported" : "Not supported"}
              </span>
            </span>
          </div>

          <div className="flex items-center gap-2">
            {isSubscribed ? (
              <FaBell className="text-blue-600" />
            ) : (
              <FaBellSlash className="text-gray-500" />
            )}
            <span className="font-medium">
              Notification Status:
              <span
                className={`ml-1 ${isSubscribed ? "text-blue-500" : "text-red-500"}`}
              >
                {isSubscribed ? "Active" : "Inactive"}
              </span>
            </span>
          </div>
        </div>

        {/* -------- action -------- */}
        <div className="flex items-center gap-2">
          {isSupported && !isSubscribed && (
            <Button
              id="enable"
              variant="blue-solid"
              size="sm"
              onClick={handleSubscribe}
            >
              {isLoadingSubscribe ? (
                <div className="h-5 w-36 flex items-center justify-center">
                  <LoaderDots />
                </div>
              ) : (
                <>
                  <FaBell className="mr-0" /> Aktifkan Notifikasi
                </>
              )}
            </Button>
          )}

          {isSubscribed && (
            <Button
              id="disabled"
              variant="red-solid"
              size="sm"
              onClick={handleUnsubscribeCurrentDevice}
            >
              {isLoadingSubscribe ? (
                <div className="h-5 w-44.25 flex items-center justify-center">
                  <LoaderDots />
                </div>
              ) : (
                <>
                  <FaBellSlash className="mr-0" />
                  Disable on This Device
                </>
              )}
            </Button>
          )}
        </div>
      </section>

      {isSubscribed && (
        <section className="mb-4 rounded-sm border border-gray-300 bg-white p-3 flex flex-col gap-2">
          <Button
            id="disabled"
            variant="blue-solid"
            size="sm"
            onClick={handleSendPushTest}
          >
            {isLoadingPushTest ? (
              <div className="h-5 w-32 flex items-center justify-center">
                <LoaderDots />
              </div>
            ) : (
              <>
                <FaBell className="mr-0" />
                Coba Notifikasi
              </>
            )}
          </Button>
        </section>
      )}

      {/* -------- device list -------- */}
      <section>
        <h2>Active Devices</h2>

        {isLoadingDevices && <p>Loading...</p>}
        {!isLoadingDevices && devices.length === 0 && (
          <p>tidak ada devices active</p>
        )}

        <div className="flex flex-col gap-2">
          {devices.map((d) => {
            const isCurrentDevice = currentDeviceId === d.device_id;
            return (
              <div
                key={d.device_id}
                className={`rounded-sm border p-3 flex flex-col ${
                  isCurrentDevice
                    ? "border-blue-500 bg-blue-50"
                    : "border-gray-300 bg-white"
                }`}
              >
                {/* ---------- header ---------- */}
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2">
                    {d.is_mobile ? (
                      <FaMobileAlt className="text-lg text-gray-600" />
                    ) : (
                      <FaDesktop className="text-lg text-gray-600" />
                    )}

                    {isEditing === d.device_id ? (
                      <TextField
                        id="device-label"
                        type="text"
                        className="text-sm px-2! py-1.5! w-30!"
                        value={labelInput}
                        onChange={(val) => setLabelInput(val as string)}
                        placeholder="Device label"
                      />
                    ) : (
                      <h3 className="font-semibold text-base">
                        {d.device_label || `${d.browser} - ${d.platform}`}
                      </h3>
                    )}
                  </div>

                  {/* ---------- edit / save ---------- */}
                  {isEditing === d.device_id ? (
                    <div className="flex gap-1">
                      <Button
                        id="save"
                        variant="green-solid"
                        size="sm"
                        onClick={() => handleRename(d.device_id)}
                      >
                        Save
                      </Button>
                      <Button
                        id="cancel"
                        variant="gray-outline"
                        size="sm"
                        onClick={() => setIsEditing(null)}
                      >
                        Cancel
                      </Button>
                    </div>
                  ) : (
                    <Button
                      id="rename"
                      variant="green-outline"
                      size="sm"
                      onClick={() => {
                        setIsEditing(d.device_id);
                        setLabelInput(d.device_label || "");
                      }}
                    >
                      <FaEdit />
                    </Button>
                  )}
                </div>

                {/* ---------- info ---------- */}
                <div className="text-sm text-gray-600">
                  {d.browser} • {d.platform} •{" "}
                  {d.is_mobile ? "Mobile" : "Desktop"}
                </div>
                <div className="flex flex-row-reverse justify-between">
                  {/* ---------- actions ---------- */}
                  <Button
                    id="logout-device"
                    variant="red-solid"
                    size="sm"
                    // disabled={isCurrentDevice}
                    onClick={() => {
                      setIsShowModalUnsubscribeOther(true);
                      setOtherDevice(d);
                    }}
                  >
                    <FaSignOutAlt className="mr-0" />
                    Logout
                  </Button>
                  {isCurrentDevice && (
                    <div className="flex items-center gap-1 text-sm text-blue-600">
                      <FaCheckCircle />
                      Device saat ini
                    </div>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      </section>

      <Modal
        isOpen={isShowModalUnsubscribeOther}
        title="Hapus notifikasi"
        confirmText="Hapus"
        cancelText="Batal"
        confirmVariant="red-solid"
        onCancel={() => {
          setIsShowModalUnsubscribeOther(false);
        }}
        onConfirm={async () => {
          if (otherDevice) {
            await handleUnsubscribeOther(otherDevice.device_id);
          }
          setIsShowModalUnsubscribeOther(false);
        }}
      >
        <div>
          Apakah anda yakin akan menghapus notifikasi dari device{" "}
          {otherDevice?.device_label ||
            `${otherDevice?.browser} - ${otherDevice?.platform}`}
          .
        </div>
        <div>
          Notifikasi hanya bisa dibuat ulang melalui device{" "}
          {otherDevice?.device_label ||
            `${otherDevice?.browser} - ${otherDevice?.platform}`}
          .
        </div>
      </Modal>

      <Modal
        isOpen={isShowModalDenied}
        title="Anda menolak notifikasi"
        confirmText="Ok"
        cancelText=""
        confirmVariant="blue-solid"
        onCancel={() => {
          setIsShowModalDenied(false);
        }}
        onConfirm={() => {
          setIsShowModalDenied(false);
        }}
      >
        <div>
          Sebelumnya anda telah menolak notifikasi. Aktifkan notifikasi melalui
          pengaturan browser agar bisa menerima notifikasi.
        </div>
        <div>
          mohon maaf tutorial belum ada, jika tidak bisa, hubungi developer
          saja.
        </div>
        <Button
          id="tutorial"
          variant="blue-solid"
          size="sm"
          disabled={true}
        >
          Tutorial cara mengaktifkan.
        </Button>
      </Modal>
    </div>
  );
}
