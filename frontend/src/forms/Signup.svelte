<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from 'svelte'
  import { get } from 'svelte/store'
  import { store } from '../ts/store'
  import { showPassword } from '../ts/showPassword'
  import { makeForm } from '../ts/makeForm'
  import { toggleRed } from '../ts/toggleRed'
  import { setRedOutline } from '../ts/setRedOutline'
  import Box from '../components/Box.svelte'
  import Loading from '../components/Loading.svelte'
  import { blankInput } from '../ts/blankInput'

  let fname: HTMLInputElement, dob: HTMLInputElement, uine: HTMLInputElement, uin: HTMLInputElement, psin: HTMLInputElement, gin: HTMLInputElement

  onMount(() => {
    fname = document.getElementById('fname') as HTMLInputElement
    dob = document.getElementById('dob') as HTMLInputElement
    uine = document.getElementById('uine') as HTMLInputElement
    uin = document.getElementById('uin') as HTMLInputElement
    psin = document.getElementById('psin') as HTMLInputElement
    gin = document.getElementById('gin') as HTMLInputElement

    store.successful_signup.update(() => false)
  })

  afterUpdate(() => blankInput())

  onDestroy(() => {
    us1()
    blankInput()
  })

  let startDate = new Date().getFullYear() - 100 + '-01-01',
    endDate = new Date().getFullYear() - 15 + '-12-31'

  afterUpdate(() => blankInput())
  onDestroy(() => window.stop())

  // Inherited from global state
  let submit_error: boolean

  let us1 = store.submit_error.subscribe(() => {
    submit_error = get(store.submit_error)
  })

  let btnFunc = async () => {
    sendAndGetOTP()
  }

  // Local state
  let w_email = false,
    w_phone = false,
    w_username = false,
    w_otp = false,
    ps_err = false,
    un_av = false,
    un_searching = false,
    uin_g_15 = false,
    fname_g_25 = false,
    otp_to_email = false,
    otp_to_phone = false,
    otp = false,
    loading = false,
    resend_otp = false,
    timeToResend = 60

  const checkUsername = async () => {
    if (uin.value.length === 0) {
      uin.removeAttribute('style')
      uin_g_15 = false
      w_username = false
      return
    } else if (uin.value.length > 15 && uin.value.trim().length > 15) {
      uin_g_15 = true
      setRedOutline(uin)
      return
    } else {
      uin_g_15 = false
      uin.removeAttribute('style')
      if (un_searching === false && uin.value.length < 16) {
        un_searching = true
        setTimeout(async () => {
          const r = await fetch(`/signup/redirect?name=${uin.value}`)
          try {
            if (!r.ok) throw Error
            else {
              const j = await r.json()
              if (j['_status'] === 'unalreg') {
                w_username = true
                un_av = false
                setRedOutline(uin)
              } else if (j['_status'] === 'done') {
                w_username = false
                un_av = true
                uin.setAttribute('style', 'outline: 1px solid #0a0; box-shadow: 0 0 0.5rem #0a0;')
              }
            }
          } catch {
            successful_signup = false
            submit_error = true
            setTimeout(() => (submit_error = false), 2000)
          } finally {
            un_searching = false
          }
        }, 2000)
      }
    }
  }

  const sendAndGetOTP = async () => {
    if (fname.value === '' || dob.value === '' || uine.value === '' || uin.value === '' || psin.value === '' || gin.value === '') toggleRed([fname, dob, gin, uine, uin, psin])
    else if (un_av && !un_searching && fname.value.length < 26 && psin.value.length < 14 && psin.value.length > 3) {
      const b = makeForm(['fname', 'dob', 'uname', 'eorp', 'pass', 'gender'], [fname.value, dob.value, uin.value, uine.value, psin.value, gin.value])
      loading = true
      const r = await fetch('/signup/redirect', { method: 'POST', body: b })
      try {
        if (!r.ok) throw Error
        else {
          const j = await r.json()
          switch (j['_status']) {
            case 'emalreg':
              w_email = true
              toggleRed([uine], () => (w_email = false))
              break
            case 'pnalreg':
              w_phone = true
              toggleRed([uine], () => (w_phone = false))
              break
            case 'otptoemail':
              otp = true
              // setTimeout(() => (otp_to_email = false), 5000)

              otp_to_email = true
              
              resend_otp = false

              let i = setInterval(() => {
                if (timeToResend === 0) {
                  timeToResend = 60
                  resend_otp = true
                  clearInterval(i)
                  return
                }
                timeToResend--
              }, 1000)
              btnFunc = async () => {
                const eotp = document.getElementById('otp') as HTMLInputElement
                if (eotp.value.length === 0) toggleRed([eotp])
                else {
                  loading = true
                  b.set('eotp', eotp.value)
                  const r = await fetch('/signup/redirect', { method: 'POST', body: b })
                  try {
                    if (!r.ok) throw Error
                    else {
                      const d = await r.json()
                      if (d['_status'] === 'success') {
                        store.successful_signup.update(() => true)
                        store.state.update(() => 'login')
                      } else if (d['_status'] === 'invalidotp') {
                        w_otp = true
                        setTimeout(() => (w_otp = false), 4000)
                      } else throw Error
                    }
                  } catch {
                    store.submit_error.update(() => true)
                  } finally {
                    loading = false
                  }
                }
              }
              break
          }
        }
      } catch {
        store.successful_signup.update(() => false)
        store.submit_error.update(() => true)
        setTimeout(() => store.submit_error.update(() => false), 2000)
      } finally {
        loading = false
      }
    }
  }
</script>

<form class="linform">
  <div class="msg">
    {#if otp_to_email}
      <Box msg={'Verify the OTP sent to ' + uine.value} green="true" />
    {:else if w_otp}
      <Box msg="Invalid OTP" />
    {:else if submit_error}
      <Box msg="An error has occured" />
    {/if}
  </div>
  {#if !otp}
    <div class="usr">
      <input
        type="text"
        class="box"
        id="fname"
        on:input={e => {
          if (e.target.value.trim().length > 25) {
            fname_g_25 = true
            setRedOutline(e.target)
            return
          }
          fname_g_25 = false
          e.target.removeAttribute('style')
        }}
        autoComplete="off"
        autoCapitalize="off"
        autoCorrect="off" />
      <label for="fname" class="usrlbl">Full Name</label>
      {#if fname_g_25}
        <div class="wstat">Full name must be under 26 characters</div>
      {/if}
    </div>
    <div class="usr">
      <input type="text" class="box" class:redbox={w_email} id="uine" autoComplete="off" autoCapitalize="off" autoCorrect="off" required />
      <label for="uine" class="usrlbl">Email address</label>
      {#if w_email}
        <div class="wstat">Email address already registered!</div>
      {:else if w_phone}
        <div class="wstat">Phone no. already registered!</div>
      {/if}
    </div>
    <div class="usr">
      <input
        type="text"
        class="box"
        class:redbox={w_username}
        id="uin"
        on:input={() => {
          checkUsername()
        }}
        autoComplete="off"
        autoCapitalize="off"
        autoCorrect="off"
        required />
      <label for="uin" class="usrlbl">Username</label>
      {#if w_username}
        <div class="wstat">Username already registered</div>
      {:else if uin_g_15}
        <div class="wstat">Username must be under 16 characters</div>
      {/if}
    </div>
    <div class="usr">
      <input
        type="password"
        class="box"
        id="psin"
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
      <label for="psin" class="usrlbl">Password</label>
      <svg on:click={() => showPassword()} class="eye" baseProfile="tiny" version="1.2" viewBox="0 0 24 24" width="1.2rem" xmlns="http://www.w3.org/2000/svg"><path d="M21.821,12.43c-0.083-0.119-2.062-2.944-4.793-4.875C15.612,6.552,13.826,6,12,6c-1.825,0-3.611,0.552-5.03,1.555  c-2.731,1.931-4.708,4.756-4.791,4.875c-0.238,0.343-0.238,0.798,0,1.141c0.083,0.119,2.06,2.944,4.791,4.875  C8.389,19.448,10.175,20,12,20c1.826,0,3.612-0.552,5.028-1.555c2.731-1.931,4.71-4.756,4.793-4.875  C22.06,13.228,22.06,12.772,21.821,12.43z M12,16.5c-1.934,0-3.5-1.57-3.5-3.5c0-1.934,1.566-3.5,3.5-3.5c1.93,0,3.5,1.566,3.5,3.5  C15.5,14.93,13.93,16.5,12,16.5z" /><g><path d="M14,13c0,1.102-0.898,2-2,2c-1.105,0-2-0.898-2-2c0-1.105,0.895-2,2-2C13.102,11,14,11.895,14,13z" /></g></svg>
      {#if ps_err}
        <div class="wstat">Password must be 4 to 14 characters</div>
      {/if}
    </div>
  {/if}
  {#if otp}
    <div class="usr">
      <input type="text" class="box" id="otp" />
      <label for="otp" class="usrlbl">Enter OTP</label>
    </div>
  {/if}
  <button
    class="btn"
    class:disabled={loading}
    disabled={loading}
    on:click={e => {
      e.preventDefault()
      btnFunc()
    }}>
    {#if loading}
      <Loading />
    {:else}
      Sign Up
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
        {#if otp}
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
      {/if}
    </ul>
  </div>
</form>

<style lang="scss">
  @import '../style/form';
  @include form;
</style>
