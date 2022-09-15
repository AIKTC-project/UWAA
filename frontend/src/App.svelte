<script lang="ts">
  import { afterUpdate } from 'svelte'
  import { get } from 'svelte/store'
  import { store } from './ts/store'
  import Login from './forms/Login.svelte'
  import Signup from './forms/Signup.svelte'
  import Forgot from './forms/Forgot.svelte'
  import { blankInput } from './ts/blankInput'

  // onMount(() => blankInput())
  afterUpdate(() => blankInput())

  // Inherited from global state
  let state: string

  store.state.subscribe(() => {
    state = get(store.state)
  })

  afterUpdate(() => {
    document.getElementsByClassName('box')[0].focus()
    store.submit_error.update(() => false)
  })
</script>

<div class="container">
  <div class="dyn">
    <div class="form" id="form">
      {#if state === 'login'}
        <Login />
      {:else if state === 'signup'}
        <Signup />
      {:else if state === 'forgot'}
        <Forgot />
      {/if}
    </div>
  </div>
  <div class="brand">
    <div class="brandinner">
      <div class="brandn">
        <span style="color: green">U</span>
        <span style="color: blue">W</span>
        <span style="color: red">A</span>
        <span style="color: brown">A</span>
        <div class="desc">We, the youth</div>
      </div>
    </div>
  </div>
</div>

<style lang="scss">
  @import './style/_vars';

  .container {
    @include flex($align: center, $justify: center);
    max-width: 15in;
    height: 100%;
    width: 100%;
    max-height: 100%;
    margin-top: auto;
    margin-bottom: auto;
  }
  .dyn {
    @include flex($align: center, $justify: center);
    width: 50%;
    height: 100%;
  }
  .form {
    padding: 0 2rem;
    border-radius: $border_radius;
    width: 20rem;
    background: transparent;
  }

  .brand {
    width: 50%;
    height: 100%;
  }
  .brandinner {
    @include flex($align: center, $justify: center, $direction: column);
    height: 100%;
    background-color: $p_color;
  }
  .brandn {
    font-size: 4rem;
    margin: 0 0 0.5rem;
    color: #fff;
  }
  .desc {
    font-size: 1.5rem;
  }

  @media (max-width: 1000px) {
    .container {
      @include flex($justify: flex-end, $direction: column-reverse);
      width: 100%;
      height: 100%;
      min-height: 4in;
    }
    .dyn {
      @include flex($align: flex-start);
      padding: 1rem 0 0;
      background-color: #fff;
      height: 85%;
      width: 100%;
    }
    .brand {
      @include flex($align: center, $justify: center);
      width: 100%;
      // min-height: 5.5rem;
      max-height: 5.5rem;
      text-align: center;
      padding: 0;
      background-color: $p_color;
    }
    .brandn {
      margin: 0;
      font-size: 2.5rem;
    }
    .desc {
      font-size: 0.9rem;
      margin-bottom: 1rem;
    }
  }
</style>
