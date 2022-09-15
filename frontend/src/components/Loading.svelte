<div class="lcon">
  <div class="loader" />
</div>

<style lang="scss">
  @import '../style/_vars';

  .lcon {
    @include flex($justify: center);
  }

  .loader {
    border: 3px solid $l_grey;
    border-top: 3px solid $p_color;
    border-radius: 50%;
    width: 0.7rem;
    height: 0.7rem;
    animation: spin 2s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
