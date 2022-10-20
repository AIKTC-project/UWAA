import { writable } from 'svelte/store'

export const store = {
	state: writable('login'),
	logged_in: writable(false),
	successful_signup: writable(false),
	submit_error: writable(false),
	pass_reset_success: writable(false),
	id: writable(''),
	username: writable(''),
	fname: writable(''),
	perks: writable(0),
}