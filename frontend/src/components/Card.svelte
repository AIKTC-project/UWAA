<script lang="ts">
import { link } from "svelte-navigator";

	export let id: string, title = '', from: string[2]
</script>

<a href={'/blog/'+id} class="con" use:link>
	<div class="title">
		{title}
	</div>
	<div class="img">
	 <img src={`/static/ballot_india/profile_img/${from[0]}.jpg`}>
	 {from[1]}
	</div>
</a>

<style lang="scss">
	@import '../style/vars';

	.con {
		@include flex($align: center, $direction: column)
		box-shadow: 0 0 0.2rem $d_grey;
		margin: 1rem;
		padding: 0.5rem;
		max-width: 15rem;
		border-radius: $border_radius;
		&:hover {
			cursor: pointer;
		}
	}

	.title {
		font-weight: bold;
		font-size: 2rem;
	}

	.img {
		@include flex($align: center);
		gap: 0.5rem;
		& > img {
		width: 2rem;
		height: 2rem;
		border-radius: 50%;
		}
	}
</style>