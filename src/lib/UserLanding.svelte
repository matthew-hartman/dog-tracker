<script>
  import { AppBar, Button, Icon, Menu, Overlay } from 'svelte-materialify';
  import { mdiMenu } from '@mdi/js';
  import DogEditList from './DogEditList.svelte';
  import DogList from './DogList.svelte';
  import { client, userID, dogs } from './stores.js';
  
  userID.subscribe(async (id) => {
    $client.realtime.subscribe("dogs" , function (e) {
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
  </AppBar>
  <DogList />
</div>
  