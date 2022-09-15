<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { get } from 'svelte/store'
  import { store } from '../ts/store'
  import { showPassword } from '../ts/showPassword'
  import { makeForm } from '../ts/makeForm'
  import { toggleRed } from '../ts/toggleRed'
  import Box from '../components/Box.svelte'
  import Loading from '../components/Loading.svelte'
import { blankInput } from '../ts/blankInput';

  onMount(() => {
    document.getElementById('cred').focus()
    setTimeout(() => store.successful_signup.update(() => false), 5000)
  })

   afterUpdate(() => blankInput())

  onDestroy(() => {
    window.stop()
    us1()
    us2()
    us3()
    blankInput()
  })

  // Inherited from global state
  let submit_error: boolean, successful_signup: boolean, pass_reset_success: boolean

  let us1 = store.submit_error.subscribe(() => {
    submit_error = get(store.submit_error)
  })

  let us2 = store.successful_signup.subscribe(() => {
    successful_signup = get(store.successful_signup)
  })

  let us3 = store.pass_reset_success.subscribe(() => {
    pass_reset_success = get(store.pass_reset_success)
  })

  // Local state
  let w_cred = false,
    ac_not_exists = false,
    loading = false

  const login = async () => {
    const cred = document.getElementById('cred') as HTMLInputElement,
      pass = document.getElementById('password') as HTMLInputElement

    if (cred.value === '' || pass.value === '') toggleRed([cred, pass])
    else {
      const b = makeForm(['cred', 'pass'], [cred.value, pass.value])
      loading = true
      const r = await fetch('/login/redirect', { method: 'POST', body: b })
      try {
        if (!r.ok) throw Error
        else {
          const j = await r.json()
          if (j['_status'] === 'oops') {
            ac_not_exists = false
            w_cred = true
            toggleRed([cred, pass], () => {
              w_cred = false
            })
          } else if (j['_status'] === 'notexists') {
            ac_not_exists = true
            toggleRed([cred], () => {}, true)
          } else if (j['_status'] === 'redirect') {
            window.location.href = '/'
          }
        }
      } catch {
        store.submit_error.update(() => true)
        setTimeout(() => store.submit_error.update(() => false), 5000)
      } finally {
        loading = false
      }
    }
  }
</script>

<form class="linform">
  <div class="msg">
    {#if successful_signup}
      <Box msg="Successfully signed up" green={true} />
    {:else if w_cred}
      <div class="w" transition:fly={{ x: -500, duration: 100 }}>
        Invalid credentials!
        <!-- svelte-ignore a11y-missing-attribute -->
        <a
          on:click={e => {
            e.preventDefault()
            store.state.update(() => 'forgot')
          }}>
          Forgot Password?
        </a>
      </div>
    {:else if pass_reset_success}
      <Box msg="Password reset successful" green="true" />
    {:else if ac_not_exists}
      <div class="w" transition:fly={{ x: -500, duration: 100 }}>
        An account with that credential does not exists!<br />
        <!-- svelte-ignore a11y-missing-attribute -->
        Just
        <a
          on:click={e => {
            e.preventDefault()
            store.state.update(() => 'signup')
          }}>
          create a new account
        </a>
        if you don't have one
      </div>
    {:else if submit_error}
      <Box msg="An error has occured!" />
    {/if}
  </div>
  <div class="usr">
    <input id="cred" type="text" class="box" autoComplete="off" autoCapitalize="off" autoCorrect="off" required />
    <label for="cred" class="usrlbl">Email or Username</label>
  </div>
  <div class="usr">
    <input id="password" type="password" class="box" required />
    <label for="password" class="usrlbl">Password</label>
    <svg on:click={() => showPassword()} class="eye" baseProfile="tiny" height="1.2rem" version="1.2" viewBox="0 0 24 24" width="1.2rem" xmlns="http://www.w3.org/2000/svg"><path d="M21.821,12.43c-0.083-0.119-2.062-2.944-4.793-4.875C15.612,6.552,13.826,6,12,6c-1.825,0-3.611,0.552-5.03,1.555  c-2.731,1.931-4.708,4.756-4.791,4.875c-0.238,0.343-0.238,0.798,0,1.141c0.083,0.119,2.06,2.944,4.791,4.875  C8.389,19.448,10.175,20,12,20c1.826,0,3.612-0.552,5.028-1.555c2.731-1.931,4.71-4.756,4.793-4.875  C22.06,13.228,22.06,12.772,21.821,12.43z M12,16.5c-1.934,0-3.5-1.57-3.5-3.5c0-1.934,1.566-3.5,3.5-3.5c1.93,0,3.5,1.566,3.5,3.5  C15.5,14.93,13.93,16.5,12,16.5z" /><g><path d="M14,13c0,1.102-0.898,2-2,2c-1.105,0-2-0.898-2-2c0-1.105,0.895-2,2-2C13.102,11,14,11.895,14,13z" /></g></svg>
  </div>
  <button
    class="btn"
    on:click={e => {
      e.preventDefault()
      login()
    }}
    class:disabled={loading}
    disabled={loading}>
    {#if loading}
      <Loading />
    {:else}
      Log In
    {/if}
  </button>
  <div class="misc">
    <!-- <hr /> -->
    <ul>
      <li>
        <button class="btn ibtn" on:click={() => store.state.update(() => 'signup')}>
          <div class="img">
            <svg style="enable-background:new 0 0 24 24; margin-top: -1" version="1.1" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><g id="info" /><g id="icons"><path d="M12,1C5.9,1,1,5.9,1,12s4.9,11,11,11s11-4.9,11-11S18.1,1,12,1z M17,14h-3v3c0,1.1-0.9,2-2,2s-2-0.9-2-2v-3H7   c-1.1,0-2-0.9-2-2c0-1.1,0.9-2,2-2h3V7c0-1.1,0.9-2,2-2s2,0.9,2,2v3h3c1.1,0,2,0.9,2,2C19,13.1,18.1,14,17,14z" id="add" /></g></svg>
          </div>
          Create new account ?</button>
      </li>
      <li>
        <button
          class="btn ibtn"
          on:click={e => {
            e.preventDefault()
            store.state.update(() => 'forgot')
          }}>
          <div class="img">
            <svg viewBox="0 0 22 22" style="margin-left: -1;" xmlns="http://www.w3.org/2000/svg"><path d="M12.26 11.74L10 14H8v2H6v2l-2 2H0v-4l8.26-8.26a6 6 0 1 1 4 4zm4.86-4.62A3 3 0 0 0 15 2a3 3 0 0 0-2.12.88l4.24 4.24z" /></svg>
          </div>
          Forgot password ?</button>
      </li>
    </ul>
  </div>
</form>

<style lang="scss">
  @import '../style/form';
  @include form;

  .w {
    padding: 0.4rem 0.2rem;
    font-size: 0.7rem;
    border-radius: $border_radius;
    text-align: center;
    color: $red;
    border: 1px solid $red;
    background-color: transparentize($red, 0.8);
  }
</style>
