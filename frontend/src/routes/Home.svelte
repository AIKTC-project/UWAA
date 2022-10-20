<script lang="ts">
    import { store } from "../ts/store";

	import Card from "../components/Card.svelte";
    import { onDestroy } from "svelte";

	const load = async () => {
		const r = await fetch("/api/blog");
		try {
			const j = await r.json();
			injectInDom(j);
		} catch {
			console.log("error");
		}
	};

	const injectInDom = (r) => {
		let ll_box = document.createElement("div");
		ll_box.className = 'bi-ll-box'

		document.getElementById("content").appendChild(ll_box);
		r.id.forEach((_, i) => {
			console.log(r.id[i])
			let comp = ll_box.appendChild(document.createElement("compbi"));

			new Card({
				target: comp,
				props: {
					id: r['id'][i],
					from: r['_from'][i],
					title: r.title[i]
				},
			});
		});
	};

	load();

	let logged_in: boolean

	const us2 = store.logged_in.subscribe((t) => (logged_in = t));
	onDestroy(() => {
		us2();
	});
</script>

<div class="con">
	<div id="content" />
</div>

<style lang="scss">
	@import "../style/_vars";

	.con {
		padding-top: 5rem;
		width: 100%;
		height: 100%;
	}

	:global(.bi-ll-box) {
		display: grid;
		align-items: center;
		width: 100%;
		grid-template-columns: 1fr 1fr 1fr 1fr;
	}	

	@media (max-width: 1000px) {
	:global(.bi-ll-box) {
		grid-template-columns: 1fr 1fr;
	}
	}
</style>
