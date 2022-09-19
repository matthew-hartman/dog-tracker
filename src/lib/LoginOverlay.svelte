<script>
  import { Button, Overlay, Card, CardTitle, CardText, List, ListItem, TextField, MaterialApp } from 'svelte-materialify';
  import { client, userID } from './stores';
  
  let type = '';

  const emailRules = [
    (v) => !!v || 'Required',
    (v) => v.length <= 25 || 'Max 25 characters',
    (v) => {
      const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return pattern.test(v) || 'Invalid e-mail.';
    },
  ];
  const passwordRules1 = [
    (v) => !!v || 'Required',
    (v) => v.length <= 25 || 'Max 25 characters',
    (v) => v.length >= 8 || 'Min 8 characters',
  ];
  const passwordRules2 = [
    ...passwordRules1,
    (v) => v == registerForm.password1.value || 'Passwords do not match',
  ];
  let registerForm = {
    email: {
      value: '',
      error: true,
    },
    password1: {
      value: '',
      error: true,
    },
    password2: {
      value: '',
      error: true,
    },
  };
  let loginForm = {
    email: {
      value: '',
      error: true,
    },
    password: {
      value: '',
      error: true,
    },
  };
  
  async function handleLogin() {
    try {
      const auth = await $client.users.authViaEmail(loginForm.email.value, loginForm.password.value);
      userID.set(auth.user.id);
    } catch (e) {
      console.log(e);
    }
  }
  
  async function handleRegister() {
    if (registerForm.email.error || registerForm.password1.error || registerForm.password2.error) {
      return;
    }
    const auth = await $client.users.create({
        email: registerForm.email.value,
        password: registerForm.password1.value,
        passwordConfirm: registerForm.password2.value,
    });
    userID.set(auth.id);
  }

</script>
  
<MaterialApp>
  <Overlay>
    <Card style="width: 25em;">
      {#if type == ''}
        <CardTitle >Login / Register</CardTitle>
        <List>
          <ListItem>
              <Button on:click={() => type = 'login'}>Login</Button>
          </ListItem>
          <ListItem>
              <Button on:click={() => type = 'register'}>Register</Button>
          </ListItem>
        </List>
      {:else if type == 'login'}
        <CardTitle>Login</CardTitle>
        <CardText>
          <TextField rules={emailRules} placeholder="Email" bind:value={loginForm.email.value} bind:error={loginForm.email.error} />
          <TextField rules={passwordRules1} placeholder="Password" type="password" bind:value={loginForm.password.value} bind:error={loginForm.password.error} />
          <Button on:click={handleLogin}>Login</Button>
        </CardText>
      {:else if type == 'register'}
        <CardTitle>Register</CardTitle>
        <CardText>
          <TextField rules={emailRules} placeholder="Email" bind:value={registerForm.email.value} bind:error={registerForm.email.error} />
          <TextField rules={passwordRules1} placeholder="Password" type="password" bind:value={registerForm.password1.value} bind:error={registerForm.password1.error} />
          <TextField rules={passwordRules2} placeholder="Confirm Password" type="password" bind:value={registerForm.password2.value} bind:error={registerForm.password2.error}/>
          <Button on:click={handleRegister}>Register</Button>
        </CardText>
      {/if}
    </Card>
  </Overlay>
</MaterialApp>