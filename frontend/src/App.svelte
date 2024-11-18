<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import {
    SelectRelacaoFile,
    SelectModelFile,
    SelectDirSaida,
    ClearFiles,
    GerarArquivos,
  } from "../wailsjs/go/main/App.js";

  let relacaoFile: string;
  let modelFile: string;
  let dirSaida: string;
  let msg: string;

  function selectRelacaoFile() {
    SelectRelacaoFile().then((result) => {
      relacaoFile = result;
    });
  }

  function selectModelFile() {
    SelectModelFile().then((result) => {
      modelFile = result;
    });
  }

  function selectDirSaida() {
    SelectDirSaida().then((result) => {
      dirSaida = result;
    });
  }

  function clearFiles() {
    ClearFiles();
    relacaoFile = "";
    modelFile = "";
    dirSaida = "";
  }

  function gerarArquivos() {
    msg = "Gerando arquivos";
    GerarArquivos().then((result) => {
      msg = result;
    });
  }
</script>

<main>
  <div>
    {#if msg}
      <p>{msg}</p>
    {/if}
  </div>
  <div class="input-box" id="input">
    <button class="btn" on:click={selectRelacaoFile}>
      Selecione o arquivo com a relação
    </button>
    {#if relacaoFile}
      <div class="result">Arquivo com a relação (*.xlsx): {relacaoFile}</div>
    {/if}
  </div>
  <div class="input-box" id="input">
    <button class="btn" on:click={selectModelFile}>
      Selecione o arquivo de modelo
    </button>
    {#if modelFile}
      <div class="result">Arquivo de modelo (*.docx): {modelFile}</div>
    {/if}
  </div>
  <div class="input-box" id="input">
    <button class="btn" on:click={selectDirSaida}>
      Selecione onde salvar
    </button>
    {#if dirSaida}
      <div class="result">Irá salvar em: {dirSaida}</div>
    {/if}
  </div>
  <div class="input-box">
    <button class="btn" on:click={() => gerarArquivos()}>
      Gerar Arquivos
    </button>
  </div>
  <div class="input-box">
    <button class="btn" on:click={clearFiles}> Limpar </button>
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
  }

  .input-box {
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
    justify-content: center;
    text-align: center;
    margin: 20px auto;
    width: 80%;
    max-width: 1024px;
  }

  .input-box button {
    display: block;
    width: 240px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box button:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }
</style>
