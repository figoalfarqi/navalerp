import { atom } from "jotai";

export const uploadFileAtom = atom<{
  folderName: string;
  isUploading: boolean;
}>({
  folderName: "",
  isUploading: false,
});
