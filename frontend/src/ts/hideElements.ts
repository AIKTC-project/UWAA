export const hideElements = (e :Array<HTMLElement>) => {
	for (let i = 0; i < e.length; i++) {
		e[i].setAttribute('style', 'display:none')
	}
}