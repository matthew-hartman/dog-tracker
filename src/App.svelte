<script>
  import LoginOverlay from './lib/LoginOverlay.svelte';
  import { client, userID } from './lib/stores.js';
  import { MaterialApp, Overlay, ProgressCircular } from 'svelte-materialify';
  import UserLanding from './lib/UserLanding.svelte';

  (async () => {
    try {
      const auth = await $client.users.refresh();
      userID.set(auth.user.id);
    } catch (e) {
      userID.set("");
    }
  })();

</script>

<div class="full">
  <MaterialApp>
    {#if $userID == null}
      <Overlay>
        <ProgressCircular color="white" indeterminate size={128} />
      </Overlay>
    {:else if $userID == ""}
      <LoginOverlay />
    {:else}
      <UserLanding />
    {/if}
  </MaterialApp>
</div>

<style>
  .full {
    height: 100%;
    width: 100%;
  }
</style>