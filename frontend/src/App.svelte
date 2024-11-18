<script lang="ts">
  import {
    SelectRelacaoFile,
    SelectModelFile,
    SelectDirSaida,
    ClearFiles,
    GerarArquivos,
  } from "../wailsjs/go/main/App.js";

  let relacaoFile: string = "";
  let modelFile: string = "";
  let dirSaida: string = "";
  let msg: string = "";
  let isLoading: boolean = false;

  let elapsedTime: number = 0;
  let timerInterval: any = null;

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
    msg = "";
    elapsedTime = 0;
    if (timerInterval) {
      clearInterval(timerInterval);
      timerInterval = null;
    }
  }

  function gerarArquivos() {
    msg = "Gerando arquivos...";
    isLoading = true;
    elapsedTime = 0;
    timerInterval = setInterval(() => {
      elapsedTime++;
    }, 1000);
    GerarArquivos().then((result) => {
      msg = result;
      isLoading = false;
      if (timerInterval) {
        clearInterval(timerInterval);
        timerInterval = null;
      }
    });
  }

  // Format the elapsed time as hh:mm:ss
  $: formattedElapsedTime = formatTime(elapsedTime);

  function formatTime(totalSeconds: number) {
    const hours = Math.floor(totalSeconds / 3600)
      .toString()
      .padStart(2, "0");
    const minutes = Math.floor((totalSeconds % 3600) / 60)
      .toString()
      .padStart(2, "0");
    const seconds = (totalSeconds % 60).toString().padStart(2, "0");
    return `${hours}:${minutes}:${seconds}`;
  }
</script>

<main>
  <h1 id="title">Gerador de processos</h1>

  {#if msg}
    <div class="message">
      {#if isLoading}
        <div class="spinner"></div>
      {/if}
      <span>{msg}</span>
      {#if isLoading}
        <span class="timer">Tempo decorrido: {formattedElapsedTime}</span>
      {/if}
    </div>
  {/if}

  <div class="container">
    <div class="card">
      <button class="btn" on:click={selectRelacaoFile} disabled={isLoading}>
        Selecione o arquivo com a relação
      </button>
      {#if relacaoFile}
        <div class="result">Arquivo com a relação (*.xlsx): {relacaoFile}</div>
      {/if}
    </div>

    <div class="card">
      <button class="btn" on:click={selectModelFile} disabled={isLoading}>
        Selecione o arquivo de modelo
      </button>
      {#if modelFile}
        <div class="result">Arquivo de modelo (*.docx): {modelFile}</div>
      {/if}
    </div>

    <div class="card">
      <button class="btn" on:click={selectDirSaida} disabled={isLoading}>
        Selecione onde salvar
      </button>
      {#if dirSaida}
        <div class="result">Irá salvar em: {dirSaida}</div>
      {/if}
    </div>

    <div class="actions">
      <button class="btn primary" on:click={gerarArquivos} disabled={isLoading}>
        {#if isLoading}
          Processando...
        {:else}
          Gerar Arquivos
        {/if}
      </button>
      <button class="btn secondary" on:click={clearFiles} disabled={isLoading}>
        Limpar
      </button>
    </div>
  </div>
</main>

<style>
  :root {
    --primary-color: #333333; /* Dark gray */
    --primary-hover-color: #222222; /* Darker gray */
    --secondary-color: #666666; /* Medium gray */
    --secondary-hover-color: #555555; /* Slightly darker gray */
    --background-color: #ffffff; /* White */
    --text-color: #000000; /* Black */
    --border-radius: 4px;
    --padding: 16px;
    --margin: 16px;
  }

  body {
    margin: 0;
    font-family: Arial, sans-serif;
    background-color: var(--background-color);
    color: var(--text-color);
  }

  #title {
    text-align: center;
    margin: 40px 0;
    font-size: 32px;
    color: var(--primary-color);
  }

  .container {
    max-width: 600px;
    margin: auto;
    padding: 0 var(--padding);
  }

  .card {
    background-color: #fff;
    padding: var(--padding);
    margin-bottom: var(--margin);
    border-radius: var(--border-radius);
  }

  .btn {
    display: inline-block;
    width: 100%;
    padding: 10px 16px;
    margin-bottom: 8px;
    font-size: 16px;
    font-weight: 600;
    text-align: center;
    text-decoration: none;
    color: #fff;
    background-color: var(--primary-color);
    border: none;
    border-radius: var(--border-radius);
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  .btn:disabled {
    opacity: 0.65;
    cursor: not-allowed;
  }

  .btn:hover:not(:disabled) {
    background-color: var(--primary-hover-color);
  }

  .btn.secondary {
    background-color: var(--secondary-color);
  }

  .btn.secondary:hover:not(:disabled) {
    background-color: var(--secondary-hover-color);
  }

  .actions {
    display: flex;
    justify-content: space-between;
    margin-top: var(--margin);
  }

  .actions .btn {
    width: auto;
    flex: 1;
    margin-right: 8px;
  }

  .actions .btn:last-child {
    margin-right: 0;
  }

  .result {
    margin-top: 8px;
    font-size: 14px;
    color: var(--secondary-color);
    word-break: break-all;
  }

  .message {
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #f0f0f0; /* Light gray */
    color: var(--text-color);
    padding: var(--padding);
    margin: var(--margin) auto;
    border: 1px solid #cccccc; /* Light gray border */
    border-radius: var(--border-radius);
    max-width: 600px;
    text-align: center;
  }

  .message .spinner {
    width: 24px;
    height: 24px;
    margin-bottom: 8px;
    border: 4px solid var(--secondary-color);
    border-top-color: var(--primary-color);
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  .message .timer {
    margin-top: 8px;
    font-size: 14px;
    color: var(--secondary-color);
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
