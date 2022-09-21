self.addEventListener('push', event => {
    console.log(event.data.json());
  
    const notification = event.data.json();
  
    const title = notification.title;
    const options = {
      body: notification.body,
      icon: notification.icon,
      vibrate: notification.vibrate,
      tag: notification.dogID,
    };
  
    self.registration.showNotification(title, options);
});
  
self.addEventListener('notificationclick', (event) => {
  fetch("/notify?dogID=" + event.notification.tag, {
    method: "POST",
  })
  event.notification.close();
});