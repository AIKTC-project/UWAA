<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from 'svelte'
  import { get } from 'svelte/store'
  import { store } from '../ts/store'
  import { makeForm } from '../ts/makeForm'
  import { toggleRed } from '../ts/toggleRed'
  import { showPassword } from '../ts/showPassword'
  import Box from '../components/Box.svelte'
  import Loading from '../components/Loading.svelte'
  import { hideElements } from '../ts/hideElements'
  import { setRedOutline } from '../ts/setRedOutline'
  import { blankInput } from '../ts/blankInput'

  let feorp: HTMLInputElement, fotp: HTMLInputElement

  onMount(() => {
    feorp = document.getElementById('feorp') as HTMLInputElement
    fotp = document.getElementById('fotp') as HTMLInputElement
  })

  afterUpdate(() => blankInput())

  onDestroy(() => {
    window.stop()
    us()
    blankInput()
  })

  let button_func = (e?: Event): any => {
    sendAndGetOTP(e)
  }

  // Inherited from global state
  let submit_error: boolean

  const us = store.submit_error.subscribe(() => {
    submit_error = get(store.submit_error)
  })

  // Local state
  let disabled = true,
    ac_exists = false,
    w_otp = false,
    otp_sent = false,
    loading = false,
    resend_otp = false,
    timeToResend = 60,
    newPass = false,
    ps_err = false

  const sendAndGetOTP = async (e?: Event) => {
    if (feorp.value === '') toggleRed([feorp])
    else {
      const b = makeForm(['feorp'], [feorp.value])
      loading = true

      const r = await fetch('/forgot', { method: 'POST', body: b })
      try {
        if (!r.ok) throw Error
        else {
          const j = await r.json()
          switch (j['_status']) {
            case 'notfound':
              ac_exists = true
              disabled = true
              toggleRed([feorp], () => (ac_exists = false))
              break
            case 'otpsent':
              otp_sent = true
              setTimeout(() => (otp_sent = false), 4000)
              resend_otp = false
              let i = setInterval(() => {
                if (timeToResend === 0) {
                  timeToResend = 60
                  clearInterval(i)
                  resend_otp = true
                  return
                }
                timeToResend--
              }, 1000)
              if (e.target !== null) {
                newPass = true
                hideElements([feorp.parentElement])
                disabled = false
                button_func = reset_pass
              }
              break
          }
        }
      } catch {
        store.submit_error.update(() => true)
        setTimeout(() => {
          store.submit_error.update(() => false)
        }, 5000)
      } finally {
        store.submit_error.update(() => false)
        loading = false
      }
    }
  }

  const reset_pass = async () => {
    const npass = document.getElementById('npass') as HTMLInputElement
    if (npass.value.length > 3 && npass.value.length !== 0 && npass.value.length < 15) {
      const b = makeForm(['feorp', 'fotp', 'npass'], [feorp.value, fotp.value, npass.value])

      const r = await fetch('/forgot', { method: 'POST', body: b })
      try {
        if (!r.ok) throw Error
        const d = await r.json()
        if (d['_status'] === 'invalidotp') {
          w_otp = true
        } else if (d['_status'] === 'resetsuccess') {
          store.pass_reset_success.update(() => true)
          store.state.update(() => 'login')
        }
      } catch {
        store.submit_error.update(() => true)
      }
    }
  }
</script>

<form class="linform">
  <div class="msg">
    {#if ac_exists}
      <Box msg="Oops! Cant find an existing account" />
    {:else if w_otp}
      <Box msg="Invalid OTP" />
    {:else if submit_error}
      <Box msg="An error has occured" />
    {:else if otp_sent}
      <Box msg="OTP sent" green="true" />
    {/if}
  </div>
  <div class="usr">
    <input id="feorp" class:disabled={!disabled} type="text" class="box" autoComplete="off" autoCapitalize="off" autoCorrect="off" required disabled={!disabled} />
    <label for="feorp" class="usrlbl">Email address</label>
  </div>
  <div class="usr">
    <input id="fotp" class:disabled type="text" class="box" autoComplete="off" autoCapitalize="off" autoCorrect="off" required {disabled} />
    <label for="fotp" class="usrlbl" class:disabled>Enter OTP</label>
  </div>
  {#if newPass}
    <div class="usr">
      <input
        type="password"
        class="box"
        id="npass"
        on:input={e => {
          if ((e.target.value.length < 4 && e.target.value.length !== 0) || e.target.value.length > 14) {
            setRedOutline(e.target)
            ps_err = true
            return
          }
          e.target.removeAttribute('style')
          ps_err = false
        }}
        required />
      <label for="npass" class="usrlbl">New Password</label>
      <svg on:click={() => showPassword()} class="eye" baseProfile="tiny" version="1.2" viewBox="0 0 24 24" width="1.2rem" xmlns="http://www.w3.org/2000/svg"><path d="M21.821,12.43c-0.083-0.119-2.062-2.944-4.793-4.875C15.612,6.552,13.826,6,12,6c-1.825,0-3.611,0.552-5.03,1.555  c-2.731,1.931-4.708,4.756-4.791,4.875c-0.238,0.343-0.238,0.798,0,1.141c0.083,0.119,2.06,2.944,4.791,4.875  C8.389,19.448,10.175,20,12,20c1.826,0,3.612-0.552,5.028-1.555c2.731-1.931,4.71-4.756,4.793-4.875  C22.06,13.228,22.06,12.772,21.821,12.43z M12,16.5c-1.934,0-3.5-1.57-3.5-3.5c0-1.934,1.566-3.5,3.5-3.5c1.93,0,3.5,1.566,3.5,3.5  C15.5,14.93,13.93,16.5,12,16.5z" /><g><path d="M14,13c0,1.102-0.898,2-2,2c-1.105,0-2-0.898-2-2c0-1.105,0.895-2,2-2C13.102,11,14,11.895,14,13z" /></g></svg>
      {#if ps_err}
        <div class="wstat">Password must be 4 to 14 characters</div>
      {/if}
    </div>
  {/if}
  <button
    class="btn"
    class:disabled={loading}
    disabled={loading}
    on:click={e => {
      e.preventDefault()
      button_func(e)
    }}>
    {#if loading}
      <Loading />
    {:else}
      Submit
    {/if}
  </button>
  <div class="misc">
    <ul>
      <li>
        <button class="btn ibtn" on:click={() => store.state.update(() => 'login')}>
          <div class="img">
            <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title /><path d="M18.875,0H4.4A1.533,1.533,0,0,0,3.125,1.573V19.048a1.74,1.74,0,0,0,1.1,1.625l8,3.2a1.75,1.75,0,0,0,2.4-1.625v-17a1.744,1.744,0,0,0-1.1-1.625L10.664,2.482A.25.25,0,0,1,10.757,2h7.868a.25.25,0,0,1,.25.25V18.725a.25.25,0,0,1-.25.25h-1.75a1,1,0,0,0,0,2h2a2,2,0,0,0,2-2V2A2,2,0,0,0,18.875,0Z" /></svg>
          </div>
          Login instead</button>
      </li>
      {#if !loading}
        <li>
          <button
            class="btn ibtn"
            class:disabled={!resend_otp}
            on:click={e => {
              e.preventDefault()
              sendAndGetOTP()
            }}
            disabled={!resend_otp}>
            <div class="img">
              <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M21,11a1,1,0,0,0-1,1,8.05,8.05,0,1,1-2.22-5.5h-2.4a1,1,0,0,0,0,2h4.53a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4.77A10,10,0,1,0,22,12,1,1,0,0,0,21,11Z" fill="#6563ff" /></svg>
            </div>
            Resend OTP
            {#if timeToResend < 60}
              {#if timeToResend != 0}
                {' in ' + timeToResend + 's'}
              {/if}
            {/if}
          </button>
        </li>
      {/if}
    </ul>
  </div>
</form>

<style lang="scss">
  @import '../style/form';
  @include form;
</style>
