import { writable } from 'svelte/store'

export const store = {
	state: writable('login'),
	successful_signup: writable(false),
	submit_error: writable(false),
	pass_reset_success: writable(false)
}