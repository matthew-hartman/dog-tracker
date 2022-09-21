<script>
  import { Container, Row, Chip, MaterialApp } from 'svelte-materialify';
  import Dogbutton from './Dogbutton.svelte'
  import { client, dogs } from './stores.js';

</script>

<MaterialApp>
  <Chip class="ma-2" size="x-large">
    Please press your dogs button to notify us you are here!
  </Chip>
  <Container >
    <Row>
      {#each $dogs as dog}
        <Dogbutton dog={dog}/>
      {/each}
    </Row>
  </Container >
</MaterialApp>