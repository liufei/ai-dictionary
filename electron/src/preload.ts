// See the Electron documentation for details on how to use preload scripts:
// https://www.electronjs.org/docs/latest/tutorial/process-model#preload-scripts
import { contextBridge, ipcRenderer } from 'electron/renderer'

contextBridge.exposeInMainWorld('ai', {
  lookup: (sentence: string, word: string) =>
    ipcRenderer.invoke('lookup', { sentence, word }),
})
