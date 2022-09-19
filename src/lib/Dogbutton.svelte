<script>
  import { MaterialApp, Button, Col } from 'svelte-materialify';
  import { afterUpdate } from 'svelte';
  import { client } from './stores.js';
  import longpress from 'longpress-svelte';

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

  function handleClick() {
    if (dog.state == "waiting") {
      $client.records.update('dogs', dog.id, {state: "notify"})
    }
  }

  function handleLong() {
    if (dog.state == "notify" || dog.state == "acked") {
      $client.records.update('dogs', dog.id, {state: "waiting"})
    }
  }
  function handleExtraLong() {
    $client.records.update('dogs', dog.id, {state: "hidden"})
  }

</script>

{#if dog.state != "hidden"}
<MaterialApp>
  <Col >
    <div use:longpress={{ duration: 200, listener: handleLong }} use:longpress={{ duration: 1000, listener: handleExtraLong }}>
      <Button size="x-large" class={classname} rounded block on:click={handleClick} >
        {dog.name}: {dog.state}: {classname}
      </Button>
    </div>
  </Col>
</MaterialApp>
{/if}