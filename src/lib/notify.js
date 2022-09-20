const vapidKey = 'BNv0koJISaKQjI_5Dy2rSThQGBi2_bWczcdf_C4V8QfwMD6ss-pC9jXNFdkAHZiKi82z2ZqmylLd-gZXH38fU9k';

export const getSubscription = async function() {
    const registration = await navigator.serviceWorker.register('worker.js', {scope: '/'});
    await navigator.serviceWorker.ready;
    return await registration.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(vapidKey)
    });
}

const urlBase64ToUint8Array = function(base64String) {
    const padding = '='.repeat((4 - base64String.length % 4) % 4);
    const base64 = (base64String + padding)
      .replace(/\-/g, '+')
      .replace(/_/g, '/');
  
    const rawData = window.atob(base64);
    const outputArray = new Uint8Array(rawData.length);
  
    for (let i = 0; i < rawData.length; ++i) {
      outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray;
}
