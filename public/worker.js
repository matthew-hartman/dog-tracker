self.addEventListener('push', event => {
    console.log('[Service Worker] Push Received.');
    console.log(`[Service Worker] Push had this data: "${event.data.text()}"`);
    console.log(event.data.json());
  
    const notification = event.data.json();
  
    const title = notification.title;
    const options = {
      body: notification.body,
      icon: notification.icon,
      vibrate: notification.vibrate,
    };
  
    event.waitUntil(self.registration.showNotification(title, options));
  });
  