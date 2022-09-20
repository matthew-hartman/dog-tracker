<script>
  import { client, userID } from './stores.js'
  import { getSubscription } from './notify.js'
  import { mdiBellCancel, mdiBellCheck  } from '@mdi/js';
  import { Chip, Icon } from 'svelte-materialify';

  let uuid = localStorage.getItem('UUID');
  let enabled = uuid ? true : false;
  console.log(uuid);

  async function toggle() {
    enabled = !enabled
    if (enabled) {
      const i = await getSubscription()
      console.log(i)
      try {
        const res = await $client.records.create('vapid', {
          subscription: i,
          user: $userID,
        })
        uuid = res.id
        localStorage.setItem('UUID', uuid)
      } catch (e) {
        console.log(e)
      }
    } else {
      try {
        await $client.records.delete('vapid', uuid)
        localStorage.removeItem('UUID')
        uuid = null
      } catch (e) {
        console.log(e)
      }
    }
  }
  async function getNotificationStatus() {
    if (uuid === null) {
      return false
    }
    const record = await $client.records.getFullList('vapid')
    for (const i of record) {
      if (i.deviceID == uuid) {
        enabled = true
      }
    }
  }
  getNotificationStatus()
</script>

<Chip>
  <button on:click={toggle}>
    <Icon path={enabled ? mdiBellCheck : mdiBellCancel} />
  </button>
</Chip>