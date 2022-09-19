<script>
  import { ButtonGroup, ButtonGroupItem, Button, Icon, List, ListItem, TextField, Chip, Divider, Dialog } from 'svelte-materialify';
  import { client, userID, dogs } from './stores.js';
  import { mdiPlusCircle } from '@mdi/js';

  let addForm = {
    name: {
      value: '',
      error: true,
    },
  }
  const nameRules = [
    (v) => !!v || 'Required',
    (v) => v.length <= 25 || 'Max 25 characters',
    (v) => v.length >= 3 || 'Min 3 characters',
    (v) => {
        for (let dog of $dogs) {
            if (dog.name == v) {
                return 'Name already taken';
            }
        }
    },
];

  async function handleAdd() {
    if (addForm.name.error) {
      return;
    }
    try {
      await $client.records.create('dogs', {
          name: addForm.name.value,
          user: $userID,
          state: "waiting",
      });
    } catch (e) {
      console.log(e);
    }
    addForm.name.value = '';
  }

  async function handleHide(dog) {
    try {
      await $client.records.update('dogs', dog.id, {
          state: "hidden",
      });
    } catch (e) {
      console.log(e);
    }
  }

  async function handleShow(dog) {
    try {
      await $client.records.update('dogs', dog.id, {
          state: "waiting",
      });
    } catch (e) {
      console.log(e);
    }
  }

  async function handleDelete(dog) {
    try {
      await $client.records.delete('dogs', dog.id);
    } catch (e) {
      console.log(e);
    }
  }
</script>

<List >
  {#each $dogs as dog}
    <ListItem>
      <Chip size="x-large">
        {dog.name}
        <Divider vertical inset class="ml-2 mr-2 grey" />
        <ButtonGroup rounded borderless>
          <ButtonGroupItem disabled={dog.state == "hidden"} on:click={() => handleHide(dog)}>Hide</ButtonGroupItem>
          <ButtonGroupItem disabled={dog.state != "hidden"} on:click={() => handleShow(dog)}>Show</ButtonGroupItem>
          <ButtonGroupItem class="red" on:click={() => handleDelete(dog)}>Delete</ButtonGroupItem>
        </ButtonGroup>
      </Chip>
    </ListItem>
  {/each}
  <ListItem>
    <TextField rules={nameRules} placeholder="Name" bind:value={addForm.name.value} bind:error={addForm.name.error} />
    <Button block class="blue white-text" on:click={handleAdd}>
      <Icon path={mdiPlusCircle} />
    </Button>
  </ListItem>
</List>