<script>
  import { Container, Row, MaterialApp } from 'svelte-materialify';
  import Dogbutton from './Dogbutton.svelte'
  import { client, dogs } from './stores.js';

</script>

<MaterialApp>
  <Container >
    <Row>
      {#each $dogs as dog}
        <Dogbutton dog={dog}/>
      {/each}
    </Row>
  </Container >
</MaterialApp>