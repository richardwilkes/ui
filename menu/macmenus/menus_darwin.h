#include <Cocoa/Cocoa.h>

typedef void *Menu;
typedef void *Item;

Item newItem(const char *title, const char *key, int modifiers);
Menu subMenu(Item item);
void setBar(Menu bar);
Menu newMenu(const char *title);
Item newSeparator();
void disposeItem(Item item);
void disposeMenu(Menu menu);
void setSubMenu(Item item, Menu subMenu);
int itemCount(Menu menu);
Item item(Menu menu, int index);
void insertItem(Menu menu, Item item, int index);
void removeItem(Menu menu, int index);
void popup(void *window, Menu menu, double x, double y, Item itemAtLocation);
void setServicesMenu(Menu menu);
void setWindowMenu(Menu menu);
void setHelpMenu(Menu menu);
