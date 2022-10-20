<script lang="ts">
  import Loading from './Loading.svelte'

  export let type = 'button'
  export let loading = false
  export let inv = false
  export let w = true // set width of button to 100%
  export let style = '' // Large style mustn't be passed here
  export let callback = () => {}
</script>

<button class="b" style={style} class:inv class:w {type} disabled={loading} on:click={() => callback()}>
  {#if loading}
    <div class="loading">
      <Loading />
    </div>
  {:else}
    <slot />
  {/if}
</button>

<style lang="scss">
  @import '../style/_vars';

  .b {
    @include flex($align:center, $justify: center);
    position: relative;
    color: #fff;
    border: none;
    border-radius: $border_radius;
    background-color: $p_color;
    padding: 0.6rem;
    font-size: 0.9rem;
    text-align: center;
    font-weight: bold;
    cursor: pointer;
    column-gap: 0.3rem;
    fill: #fff;
    &:hover {
      background-color: $d_focus_color;
    }
  }

  .inv {
    background-color: transparent;
    border: 0.15rem solid $p_color;
    padding: 0.45rem;
    color: $p_color;
    fill: $p_color;
    &:hover {
      background-color: $l_hover_color;
    }
  }

  .w {
    width: 100%;
  }
</style>
