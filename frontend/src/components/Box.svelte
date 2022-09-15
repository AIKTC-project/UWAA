<script lang="ts">
	import {fly} from 'svelte/transition'

	export let msg: String
	export let green = false
</script>

<div class="w" class:green transition:fly={{ x: -500, duration: 100 }}>{@html msg}</div>

<style lang="scss">
@import '../style/_vars';

.w {
    padding: 0.4rem 0.2rem;
    font-size: 0.75rem;
    border-radius: $border_radius;
    text-align: center;
    color: $red;
    border: 1px solid $red;
    background-color: transparentize($red, 0.8);
  }

  .green {
  	color: $green;
    border: 1px solid $green;
    background-color: transparentize($green, 0.8);
  }
</style>