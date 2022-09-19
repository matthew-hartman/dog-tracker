<script>
  import { client } from './stores.js'
  import { getSubscription } from './notify.js'

  let enabled = false
  async function toggle() {
    enabled = !enabled
    if (enabled) {
      const i = await getSubscription()
      const user = await $client.users.refresh();

      const r = await $client.records.create('vapid', {
        subscription: i,
        user: user.user.id,
      })
    }
  }
  async function getNotificationStatus() {
    const record = await $client.records.getFullList('vapid')
    if (record.length > 0) {
      enabled = true
    }
  }
  getNotificationStatus()
</script>

<main>
  <button on:click={toggle}>
    {enabled ? 'Disable' : 'Enable'} notifications
  </button>
</main>