% min(X, Y, Z) - Defines a minimum function that calculates the minimum of two numbers X and Y, and returns the result in Z.
min(X, Y, Z) :- Z is min(X, Y).

% dataset(DirectoryName) - Specifies the directory where the image dataset is located.
dataset('C:\\Users\\linfe\\OneDrive\\Desktop\\csi2120\\project\\part1\\imageDataset2_15_20').

% directory_textfiles(DirectoryName, ListOfTextfiles) - Produces the list of text files in a given directory.
directory_textfiles(D, Textfiles) :-
    directory_files(D, Files),
    include(isTextFile, Files, Textfiles).

% isTextFile(Filename) - Checks if a given filename ends with '.txt', indicating it is a text file.
isTextFile(Filename) :-
    string_concat(_, '.txt', Filename).

% read_hist_file(Filename, ListOfNumbers) - Reads a histogram file and produces a list of numbers (bin values).
read_hist_file(Filename, Numbers) :-
    open(Filename, read, Stream),
    read_line_to_string(Stream, _),
    read_line_to_string(Stream, String),
    close(Stream),
    atomic_list_concat(List, ' ', String),
    atoms_numbers(List, Numbers).

% similarity_search(QueryFile, SimilarImageList) - Returns the list of images similar to the query image.
% Similar images are specified as (ImageName, SimilarityScore).
% Predicate dataset/1 provides the location of the image set.
similarity_search(QueryFile, SimilarList) :-
    dataset(D),
    directory_textfiles(D, TxtFiles),
    similarity_search(QueryFile, D, TxtFiles, SimilarList).

% similarity_search(QueryFile, DatasetDirectory, HistoFileList, SimilarImageList) - Finds similar images to the query image.
similarity_search(QueryFile, DatasetDirectory, DatasetFiles, Best) :-
    read_hist_file(QueryFile, QueryHisto),
    compare_histograms(QueryHisto, DatasetDirectory, DatasetFiles, Scores),
    sort(2, @>, Scores, Sorted),
    take(Sorted, 5, Best).

% compare_histograms(QueryHisto, DatasetDirectory, DatasetFiles, Scores) - Compares a query histogram with a list of histogram files.
compare_histograms(QueryHisto, DatasetDirectory, DatasetFiles, Scores) :-
    findall((File, Score),
            (member(File, DatasetFiles),
             atom_concat(DatasetDirectory, '/', DirWithSlash),
             atom_concat(DirWithSlash, File, Path),
             read_hist_file(Path, Histo),
             histogram_intersection(QueryHisto, Histo, Score)),
            Scores).

% histogram_intersection(Histogram1, Histogram2, Score) - Computes the intersection similarity score between two histograms.
% Score is between 0.0 and 1.0 (1.0 for identical histograms).
histogram_intersection(H1, H2, Score) :-
    maplist(min, H1, H2, MinList),
    sumlist(MinList, TotalMin),
    sumlist(H1, TotalH1),
    Score is TotalMin / TotalH1.

% take(List, K, KList) - Extracts the K first items in a list.
take(Src, N, L) :-
    findall(E, (nth1(I, Src, E), I =< N), L).

% atoms_numbers(ListOfAtoms, ListOfNumbers) - Converts a list of atoms into a list of numbers.
atoms_numbers([], []).
atoms_numbers([X|L], [Y|T]) :-
    atom_number(X, Y),
    atoms_numbers(L, T).
atoms_numbers([X|L], T) :-
    \+atom_number(X, _),
    atoms_numbers(L, T).
