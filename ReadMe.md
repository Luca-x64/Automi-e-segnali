# Automi e segnali

Progetto sviluppato per il **Laboratorio di Algoritmi e Strutture Dati**.

Questa repository raccoglie l'implementazione iniziale del progetto **"Automi e segnali"**, nata come elaborato universitario e mantenuta come base di partenza per evoluzioni successive. La release `v1.0` rappresenta una fotografia della versione originale del progetto, su cui verranno poi introdotti miglioramenti, ottimizzazioni e nuove funzionalità.

## Contesto

Il programma modella un piano cartesiano in cui possono essere presenti:

- **automi**, identificati da un nome binario e da una posizione;
- **ostacoli** rettangolari, che occupano aree non attraversabili del piano;
- **segnali di richiamo**, che permettono di verificare se alcuni automi possono raggiungere un punto del piano.

L'obiettivo principale del progetto è gestire queste entità e supportare le operazioni richieste dalla traccia, mantenendo una struttura dati coerente con il problema.

## Funzionalità principali

Il programma supporta le seguenti operazioni:

- creazione o reset del piano;
- inserimento e riposizionamento di automi;
- inserimento di ostacoli rettangolari;
- interrogazione dello stato di una coordinata;
- stampa del contenuto del piano;
- ricerca delle posizioni degli automi dato un prefisso del nome;
- verifica dell'esistenza di un percorso libero di lunghezza esattamente pari alla distanza di Manhattan;
- richiamo degli automi raggiungibili con un certo prefisso;
- chiusura del programma.

## Scelte implementative

### Rappresentazione degli automi

Gli automi sono memorizzati tramite un **albero binario**:

- ogni arco verso sinistra rappresenta il bit `0`;
- ogni arco verso destra rappresenta il bit `1`;
- il percorso dalla radice a un nodo corrisponde al nome binario dell'automa.

Questa scelta rende naturale:

- l'inserimento di un automa in base al suo nome;
- la ricerca di un automa esatto;
- la ricerca di tutti gli automi che condividono un prefisso.

### Rappresentazione degli ostacoli

Gli ostacoli sono memorizzati come rettangoli definiti da due coordinate:

- `p0`: vertice in basso a sinistra;
- `p1`: vertice in alto a destra.

Una coordinata del piano è considerata libera se non è contenuta in alcun ostacolo.

### Percorsi e distanza

Per verificare la raggiungibilità di un punto si considera la **distanza di Manhattan**:

`D(A, B) = |x_B - x_A| + |y_B - y_A|`

L'operazione di verifica del percorso controlla l'esistenza di un cammino libero la cui lunghezza sia **esattamente uguale** a questa distanza.

L'implementazione distingue due casi:

1. **segmento orizzontale o verticale**: viene effettuato un controllo lineare del tratto tra i due punti;
2. **coordinate diverse su entrambi gli assi**: viene eseguita una visita del piano limitata alle sole mosse che avvicinano alla destinazione, evitando percorsi che produrrebbero una lunghezza maggiore della distanza di Manhattan.

## Struttura del progetto

- `31974A_Ghirimoldi_Luca.go` — file principale, parsing dei comandi e avvio del programma.
- `utils/piano.go` — definizione dei tipi principali e delle strutture dati.
- `utils/functions.go` — implementazione delle operazioni richieste dalla traccia.
- `utils/iteratore.go` — iteratore sugli automi contenuti in un sottoalbero.
- `utils/coda.go` — coda generica di supporto.
- `utils/cordinata.go` — tipo coordinata e distanza di Manhattan.
- `formato_test.go`, `lib_test.go`, `utils_test.go` — suite di test.
- `relazione.pdf` — relazione estesa della versione originale del progetto.

## Complessità

Di seguito una panoramica delle principali complessità dell'implementazione corrente.

### Albero degli automi

Sia `L` la lunghezza del nome binario di un automa.

- **Ricerca di un automa/prefisso**: `Θ(L)`.
- **Inserimento di un automa**: `Θ(L)` nel caso peggiore.

Più precisamente, se `M` è il numero di nuovi nodi da creare:

- caso ottimale: `O(1)` se il percorso è già presente e va solo aggiornato il nodo finale;
- caso peggiore: `Θ(L)` se il nome richiede la creazione di tutti i nodi del cammino;
- caso intermedio: `Θ(M)` con `M <= L`.

### Iterazione su un sottoalbero

Sia `N` il numero di nodi del sottoalbero visitato.

- costruzione dell'iteratore: `O(N)` tempo;
- spazio occupato dalla struttura di supporto: `O(N)`;
- `HasNext()` e `Next()`: `O(1)`.

### Coda generica

Per una coda con `N` elementi:

- `push`: `O(1)` ammortizzato;
- `pop`: `O(1)`;
- `isEmpty`: `O(1)`;
- spazio: `O(N)`.

### Controllo di coordinate libere

Sia `K` il numero di ostacoli presenti nel piano.

- `coordinateLibere(x, y)`: `O(K)`.

### Verifica dell'esistenza di un percorso libero

Il costo dipende dal caso considerato:

- se il punto è interno a un ostacolo: controllo immediato dopo la scansione degli ostacoli;
- se il percorso è orizzontale o verticale: controllo lineare sul segmento, con verifica delle coordinate attraversate;
- nel caso generale: visita delle sole coordinate che possono appartenere a un percorso minimo rispetto alla distanza di Manhattan.

In questa implementazione, il costo dipende quindi sia dal numero di ostacoli sia dal numero di coordinate esplorate durante la visita.

### Richiamo

L'operazione di richiamo:

1. individua il nodo associato al prefisso;
2. visita gli automi del sottoalbero;
3. verifica per ciascuno l'esistenza di un percorso libero;
4. seleziona quelli a distanza minima.

Di conseguenza, il suo costo dipende dal numero di automi compatibili col prefisso e dal costo delle singole verifiche di raggiungibilità.

## Esecuzione

Per eseguire il programma:

```bash
go run .
```

oppure:

```bash
go run 31974A_Ghirimoldi_Luca.go
```

## Test

Per eseguire i test principali:

```bash
go test -run TestBase -v -count=1
go test -run TestFormatoStato -v -count=1
```

## Roadmap

Questa `v1.0` rappresenta la versione iniziale del progetto. Nelle versioni successive l'obiettivo è:

- migliorare e ottimizzare alcune parti dell'implementazione;
- estendere la documentazione direttamente nel repository;
- ridurre progressivamente il ruolo della relazione estesa, mantenendo nel README solo le informazioni essenziali;
- evolvere il progetto oltre il contesto originale dell'esame.
