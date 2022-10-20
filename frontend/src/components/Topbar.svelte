<script lang="ts">
	import { get } from "svelte/store";
	import { store } from "../ts/store";
	import { link } from "svelte-navigator";
	import { onDestroy } from "svelte";
	import Button from "./Button.svelte";

	onDestroy(() => {
		us1();
		us2();
		us3();
	});

	let perks: number, fname: string, logged_in: boolean;

	// const path = useLocation()
	// const navigate = useNavigate()

	// $: pathname = $path.pathname

	const us1 = store.perks.subscribe((t) => (perks = t));

	const us2 = store.logged_in.subscribe((t) => (logged_in = t));

	const us3 = store.fname.subscribe((t) => (fname = t));
</script>

<div id="topbar" class="topbar">
	<div class="user">
		{#if logged_in}
			<a href={"https://127.0.0.1/profile/" + get(store.id)} class="act">
				<img
					class="img"
					src={`/static/ballot_india/profile_img/${get(store.id)}.jpg`}
					alt=""
				/>
				<div class="inuser">
					<div class="name">
						{fname}
					</div>
					<div class="perks" title={perks.toString()}>
						<b>{perks}</b> perks
					</div>
				</div>
			</a>
		{:else}
			<Button
				style={"margin-left:1rem"}
				inv={true}
				callback={() => (window.location.href = "/accounts")}
			>
				<div class="ico">
					
				</div>
				Login / Sign up
			</Button>
		{/if}
	</div>

	<div class="brand">
		<!-- <svg on:click={() => (menu = false)} id="menuimg" class="menuimg" style="background-color: #d6d6d6" width="2rem" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg"><title /><g id="Fill"><polygon points="28.71 4.71 27.29 3.29 16 14.59 4.71 3.29 3.29 4.71 14.59 16 3.29 27.29 4.71 28.71 16 17.41 27.29 28.71 28.71 27.29 17.41 16 28.71 4.71" /></g></svg> -->
		<a href="/" on:click={() => (window.location.pathname = "/")}>
			<span style="color: green">U</span>
			<span style="color: blue">W</span>
			<span style="color: red">A</span>
			<span style="color: brown">A</span>
		</a>
	</div>
</div>

<style lang="scss" scoped>
	@import "../style/_vars.scss";

	a {
		&:hover {
			text-decoration: none;
		}
	}

	.img {
		width: 2.5rem;
		height: 2.5rem;
		border-radius: 50%;
		margin-right: 1rem;
	}

	.topbar {
		@include flex($align: center, $justify: space-between);
		position: absolute;
		background-color: #fff;
		box-shadow: 0 0 0.1rem $d_grey;
		width: 100%;
		height: 3rem;
		padding: 0;
		z-index: 11;
		top: 0;
	}

	.brand {
		@include flex($align: center, $justify: center);
		font-size: 1.5rem;
		font-weight: bold;
		user-select: none;
		margin-right: 0.7rem;
		& > a:hover {
			text-decoration: none;
		}
	}

	.user {
		@include flex($align: center, $justify: center);
	}

	.act {
		@include flex($align: center);
		margin-left: 1rem;
	}

	.ico {
		fill: $p_color;
		width: 1rem;
	}

	.inuser {
		@include flex($direction: column);
		margin-right: 1rem;
	}

	.name {
		font-weight: bold;
	}

	.perks {
		font-size: 0.9rem;
	}
</style>
