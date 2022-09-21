<script>
  import { AppBar, Button, Icon, Menu, Overlay } from 'svelte-materialify';
  import { mdiMenu } from '@mdi/js';
  import DogEditList from './DogEditList.svelte';
  import DogList from './DogList.svelte';
  import { client, userID, dogs } from './stores.js';
  import Notification from './Notification.svelte';
  import { toast, SvelteToast } from '@zerodevx/svelte-toast'
  
  userID.subscribe(async (id) => {
    $client.realtime.subscribe("dogs" , function (e) {
      if (e.action == "update" && e.record.state == "acked") {
          toast.push( `${e.record.name} is on their way`, {dismissable: false} )
      }

      let tmp = $dogs
      let found = false
      for (const i of $dogs.keys()) {
        if (tmp[i].id == e.record.id) {
          if (e.action == "delete") {
            tmp.splice(i, 1)
          } else {
            tmp[i] = e.record
          }
          found = true
        }
      }
      if (!found) {
        tmp.push(e.record)
      }
      dogs.set(tmp)
    });
    dogs.set(await $client.records.getFullList('dogs'))
  });
  </script>
  
<div>
  <AppBar>
    <div slot="icon">
      <Menu closeOnClick={false} >
      <Button slot="activator" fab depressed >
        <Icon path={mdiMenu} />
      </Button>
      <DogEditList />
      </Menu>
    </div>
    <span slot="title">Zoomies Notifications</span>
    <div style="flex-grow:1" />
    <Notification />
  </AppBar>
  <DogList />
  <SvelteToast options={{ reversed: true, intro: { y: 192 } }} />
</div>

<style>
  :root {
    --toastContainerTop: auto;
    --toastContainerRight: auto;
    --toastContainerBottom: 8rem;
    --toastContainerLeft: calc(50vw - 8rem);
  }
</style>
