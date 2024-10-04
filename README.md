# Ascii-art

## Description

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:

```
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```

## Usage

To use this program you need to clone this repo : 
```
git clone https://zone01normandie.org/git/dbourcie/ascii-art.git
```
Then you have to run it by executing this command line : 
```
go run . <Your text here> <style>
```
You can chose the *style* of the output between *standard*, *shadow* and *thinkertoy*. By default, if you don't write any style it will output in *standard* style.

What looks like the different styles :
```
 Standard          Shadow          Thinkertoy

    /\              _|_|              O   
   /  \           _|    _|           / \  
  / /\ \          _|_|_|_|          o---o 
 / ____ \         _|    _|          |   | 
/_/    \_\        _|    _|          o   o 
 ```                 

And it will output your string in a graphic representation using ASCII

## Examples

```
student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$
student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
student$ go run . "hello" | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
student$ go run . "Hello There" | cat -e
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $
student$ go run . "1Hello 2There" | cat -e
     _    _          _   _                      _______   _                           $
 _  | |  | |        | | | |              ____  |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___       |___ \    | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \        __) |   | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |      / __/    | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/      |_____|   |_|    |_| |_|  \___| |_|     \___| $
                                                                                      $
                                                                                      $
student$ go run . "{Hello There}" | cat -e
   __  _    _          _   _               _______   _                           __    $
  / / | |  | |        | | | |             |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                            /_/   $
                                                                                       $
student$ go run . "Hello\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
student$ go run . "Hello\n\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
student$ go run . "Hello There" shadow | cat -e
                                                                                      $
_|    _|          _| _|                _|_|_|_|_| _|                                  $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|       $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| $
                                                                                      $
                                                                                      $
students$ go run . "Hello There" thinkertoy | cat -e
                                              $
o  o     o o           o-O-o o                $
|  |     | |             |   |                $
O--O o-o | | o-o         |   O--o o-o o-o o-o $
|  | |-' | | | |         |   |  | |-' |   |-' $
o  o o-o o o o-o         o   o  o o-o o   o-o $
                                              $
                                              $
students$

```
