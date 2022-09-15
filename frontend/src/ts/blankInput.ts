export const blankInput = () => {
	const b = document.getElementsByClassName('box')

	for (let i = 0; i < b.length; i++) {
		b[i].addEventListener('input', () => {
			if (b[i].value == "") {
				b[i].parentElement.children[1].classList.remove('usrlbl_focused')
				return
			}
			b[i].parentElement.children[1].classList.add('usrlbl_focused')
		})
	}
}