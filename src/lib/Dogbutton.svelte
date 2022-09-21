<script>
  import { MaterialApp, Button, Col } from 'svelte-materialify';
  import { afterUpdate } from 'svelte';
  import { client } from './stores.js';

  export let dog;

  let classname = "primary-color";
  afterUpdate(() => {
    if (dog.state == "waiting") {
      classname = "secondary-color"
    } else if (dog.state == "notify") {
      classname = "deep-purple accent-1"
    } else if (dog.state == "acked") {
      classname = "deep-purple accent-2"
    } else {
      classname = "primary-color"
    }
	});

  function handleClick(e) {
    if (e.timeStamp - cooldown < 5500) {
      return
    }
    if (dog.state == "waiting") {
      $client.records.update('dogs', dog.id, {state: "notify"})
    }
  }

  let timer;
  let cooldown;

  function updateDog(e) {
    if (dog.state == "notify" || dog.state == "acked") {
      $client.records.update('dogs', dog.id, {state: "waiting"})
    }
    if (dog.state == "waiting") {
      $client.records.update('dogs', dog.id, {state: "hidden"})
    }
    cooldown = e.timeStamp
  }

  function handleDown(e) {
    timer = setTimeout(() => {
      updateDog(e)
    }, 500)
  }
  function handleUp(e) {
    clearTimeout(timer);
  }

</script>

{#if dog.state != "hidden"}
<MaterialApp>
  <Col >
    <div on:touchstart={handleDown} on:touchend={handleUp} on:mousedown={handleDown} on:mouseup={handleUp}>
      <Button size="x-large" class={classname} rounded block on:click={handleClick} >
        {dog.name}
      </Button>
    </div>
  </Col>
</MaterialApp>
{/if}