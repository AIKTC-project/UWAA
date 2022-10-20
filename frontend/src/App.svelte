<script lang="ts">
import { onMount, afterUpdate, onDestroy } from "svelte";
import { get } from "svelte/store";
import { store } from "./ts/store";
import Login from "./forms/Login.svelte";
import Signup from "./forms/Signup.svelte";
import Forgot from "./forms/Forgot.svelte";
import { blankInput } from "./ts/blankInput";
import { Router, Route, link } from "svelte-navigator";
import Home from "./routes/Home.svelte";
import Topbar from "./components/Topbar.svelte";
    import Blog from "./routes/Blog.svelte";
    import Create from "./routes/Create.svelte";

onMount(async () => {
  const r = await fetch('/api/accounts/profile?q=self')
  try {
    const j = JSON.parse(await r.text())
    if (j['_status'] === 'ltc') throw Error()
      store.id.update(() => j['id'])
    store.username.update(() => j['uname'])
    store.fname.update(() => j['fname'])
    store.perks.update(() => j['perks'])
    store.logged_in.update(() => true)
  } catch (error) {
    store.logged_in.update(() => false)
  }
})

  // Inherited from global state
  let state: string;

  store.state.subscribe(() => {
    state = get(store.state);
  });

  let logged_in: boolean

  const us2 = store.logged_in.subscribe((t) => (logged_in = t));
  onDestroy(() => {
    us2();
  });

  /*afterUpdate(() => {
    document.getElementsByClassName("box")[0].focus();
    store.submit_error.update(() => false);
  });*/
</script>

<Router>
  {#if window.location.pathname !== '/accounts'}
  <Topbar />
  {/if}
  <div class="container">
    <Route path="/">
      <Home />
    </Route>
    <Route path="/blog/:id" let:params><Blog id={params.id} /></Route>
    <Route path="/create"><Create /></Route>
    <Route path="/accounts">
      <div class="dyn">
        <div class="form" id="form">
          {#if state === "login"}
          <Login />
          {:else if state === "signup"}
          <Signup />
          {:else if state === "forgot"}
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
    </Route>
    {#if logged_in && window.location.pathname !== '/create'}
      <a class="create" href="/create" use:link>
        <svg width="3rem" viewBox="0 0 24 24"><path d="M12.3023235,7.94519388 L4.69610276,15.549589 C4.29095108,15.9079238 4.04030835,16.4092335 4,16.8678295 L4,20.0029438 L7.06398288,20.004826 C7.5982069,19.9670062 8.09548693,19.7183782 8.49479322,19.2616227 L16.0567001,11.6997158 L12.3023235,7.94519388 Z M13.7167068,6.53115006 L17.4709137,10.2855022 L19.8647941,7.89162181 C19.9513987,7.80501747 20.0000526,7.68755666 20.0000526,7.56507948 C20.0000526,7.4426023 19.9513987,7.32514149 19.8647932,7.23853626 L16.7611243,4.13485646 C16.6754884,4.04854589 16.5589355,4 16.43735,4 C16.3157645,4 16.1992116,4.04854589 16.1135757,4.13485646 L13.7167068,6.53115006 Z M16.43735,2 C17.0920882,2 17.7197259,2.26141978 18.1781068,2.7234227 L21.2790059,5.82432181 C21.7406843,6.28599904 22.0000526,6.91216845 22.0000526,7.56507948 C22.0000526,8.21799052 21.7406843,8.84415992 21.2790068,9.30583626 L9.95750718,20.6237545 C9.25902448,21.4294925 8.26890003,21.9245308 7.1346,22.0023295 L2,22.0023295 L2,21.0023295 L2.00324765,16.7873015 C2.08843822,15.7328366 2.57866679,14.7523321 3.32649633,14.0934196 L14.6953877,2.72462818 C15.1563921,2.2608295 15.7833514,2 16.43735,2 Z" fill-rule="evenodd" /></svg>
      </a>
    {/if}
  </div>
</Router>

<style lang="scss">
@import "./style/_vars";

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

.create {
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  width: 4rem;
  height: 4rem;
  background: $p_color;
  bottom: 1rem;
  right: 1rem;
  border-radius: 50%;
  & > svg {
    fill: #fff;
  }
}

.brand {
  width: 50%;
  height: 100%;
}
.brandinner {
  @include flex($align: center, $justify: center, $direction: column);
  height: 100%;
}
.brandn {
  font-size: 4rem;
  margin: 0 0 0.5rem;
  color: #fff;
}
.desc {
  font-size: 1.5rem;
  color: $p_color;
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
