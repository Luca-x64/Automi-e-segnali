# Automi e segnali

Progetto sviluppato per il corso di **Algoritmi e Strutture Dati**.

## Cos’è il progetto

`Automi e segnali` è un simulatore su piano cartesiano che gestisce:

- **automi** (nome binario + posizione);
- **ostacoli rettangolari**;
- **segnali di richiamo** per selezionare gli automi raggiungibili.

L’obiettivo è modellare un problema reale di instradamento/richiamo su griglia, mostrando competenze su progettazione di strutture dati, gestione input testuale e implementazione di algoritmi di visita su spazio discreto.

## Cosa fa e come lo fa

Il programma legge comandi da standard input e supporta le principali operazioni previste dalla traccia:

- creazione/reset del piano;
- inserimento e riposizionamento automi;
- inserimento ostacoli;
- interrogazioni sullo stato del piano;
- ricerca per prefisso dei nomi binari;
- verifica di raggiungibilità;
- richiamo degli automi “migliori” rispetto al segnale.

L’implementazione corrente (release candidata **v1.1.0**) include il passaggio da logica basata su **BFS** a **DFS** per la visita del piano e un miglioramento della strategia di scelta degli automi da richiamare.

## Algoritmi e strutture dati usate (focus corso)

### 1) Albero binario per gli automi
Gli automi sono indicizzati in un **albero binario sui bit del nome** (`0` a sinistra, `1` a destra).

**Perché questa scelta:** ricerca naturale per nome esatto, ricerca per prefisso e coerenza con identificativi binari.

### 2) Collezione di rettangoli per gli ostacoli
Gli ostacoli sono memorizzati come rettangoli definiti da due vertici.

**Perché questa scelta:** rappresentazione diretta del dominio, controlli semplici sulle coordinate e buona estendibilità.

### 3) Visita su griglia e distanza di Manhattan
La raggiungibilità viene valutata sul piano discreto con vincolo di cammini compatibili con la distanza di Manhattan.

**Perché questa scelta:** aderenza alla traccia, separazione chiara tra logica geometrica e logica di visita, base per ulteriori ottimizzazioni.

<img width="741" height="448" alt="image" src="https://github.com/user-attachments/assets/32e2e07d-8bbd-4958-90ae-78b3230dc06c" />


## Struttura del repository

- `31974A_Ghirimoldi_Luca.go`: entrypoint e parsing comandi.
- `utils/`: logica applicativa e strutture di supporto.
- `test/input/`: file di input per i casi di test.
- `test/output/`: output attesi corrispondenti.
- `relazione.pdf`: documento di approfondimento.
- `traccia.pdf` *(consigliato, se pubblicabile)*: testo ufficiale del progetto, utile a chi vuole valutare requisiti e vincoli originali.

## Esecuzione rapida

```bash
go run .
```

## Test (3 casi principali)

Per tenere più pulita la root, i file di test sono organizzati nella cartella `test/`.
Nella relazione sono riportati i 3 comandi per verificare i casi mostrati:

```bash
go run . < test/input/input1.txt | diff -u - test/output/out1.txt
go run . < test/input/input2.txt | diff -u - test/output/out2.txt
go run . < test/input/input3.txt | diff -u - test/output/out3.txt
```
<img width="936" height="894" alt="image" src="https://github.com/user-attachments/assets/586effd0-aafd-4a42-92a6-377321e18108" />
(test n° 3)

## Nota sulla relazione

La `relazione.pdf` contiene una spiegazione più approfondita, molte scelte implementative e il razionale progettuale completo. Fa riferimento alla release **v1.0.0** (versione con BFS), quindi è da considerare come base storica rispetto all’evoluzione corrente.
